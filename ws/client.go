// Package ws provides a WebSocket client for receiving real-time sync
// notifications
package ws

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

const (
	baseURL   = "wss://ws.todoist.com/ws"
	originURL = "https://app.todoist.com"
)

const (
	pingInterval = 60 * time.Second
	maxBackoff   = 128 * time.Second
)

type Client struct {
	token   string
	handler Handler

	mu   *sync.Mutex
	conn *websocket.Conn

	wg     *sync.WaitGroup
	once   sync.Once
	cancel context.CancelFunc
	msgCh  chan Message
}

func NewClient(token string, handler Handler) *Client {
	return &Client{
		token:   token,
		handler: handler,
		mu:      &sync.Mutex{},
		wg:      &sync.WaitGroup{},
		once:    sync.Once{},
		msgCh:   make(chan Message),
	}
}

func (c *Client) Listen(ctx context.Context) {
	c.once.Do(func() {
		ctx, c.cancel = context.WithCancel(ctx)

		c.wg.Add(1)
		go c.handleMessage(ctx)

		select {
		case c.msgCh <- Disconnected:
		case <-ctx.Done():
			return
		}
	})
}

func (c *Client) Close() error {
	if c.cancel != nil {
		c.cancel()
	}

	c.wg.Wait()
	c.closeConn()

	return nil
}

func (c *Client) dial(ctx context.Context) (*websocket.Conn, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u.RawQuery = "token=" + c.token

	header := make(http.Header)
	header.Add("Origin", originURL)
	opts := &websocket.DialOptions{HTTPHeader: header}

	conn, _, err := websocket.Dial(ctx, u.String(), opts)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) listen(ctx context.Context) {
	defer func() {
		c.wg.Done()
		c.closeConn()

		select {
		case c.msgCh <- Disconnected:
		case <-ctx.Done():
			return
		}
	}()

	timeout := time.NewTimer(pingInterval)
	defer timeout.Stop()

	c.msgCh <- Connected

	for {
		select {
		case <-ctx.Done():
			return
		case <-timeout.C:
			return
		default:
			c.mu.Lock()
			conn := c.conn
			c.mu.Unlock()

			if conn == nil {
				return
			}

			msg := &message{}
			if err := wsjson.Read(ctx, conn, msg); err != nil {
				continue
			}

			if msg.Type == Ping {
				log.Print(msg.Type)
				timeout.Reset(pingInterval)
			}

			select {
			case c.msgCh <- msg.Type:
			case <-ctx.Done():
				return
			}
		}
	}
}

func (c *Client) handleMessage(ctx context.Context) {
	defer c.wg.Done()
	backoff := newBackoff(1*time.Second, maxBackoff)

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-c.msgCh:
			switch msg {
			case Disconnected:
				for {
					if conn, err := c.dial(ctx); err == nil {
						c.mu.Lock()
						c.conn = conn
						c.mu.Unlock()

						backoff.Reset()
						c.wg.Add(1)
						go c.listen(ctx)
						break
					}

					d := backoff.Duration()
					log.Printf("failed to connect, next retry in %v", d)

					select {
					case <-ctx.Done():
						return
					case <-time.After(d):
						continue
					}
				}
				fallthrough
			default:
				if err := c.handler.HandleMessage(ctx, msg); err != nil {
					log.Print("failed to handle message:", err)
				}
			}
		}
	}
}

func (c *Client) closeConn() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		c.conn.Close(websocket.StatusNormalClosure, "")
		c.conn = nil
	}
}
