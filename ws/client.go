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

	done  chan struct{}
	msgCh chan Message
}

func NewClient(token string, handler Handler) *Client {
	return &Client{
		token:   token,
		handler: handler,
		mu:      &sync.Mutex{},
		done:    make(chan struct{}),
		msgCh:   make(chan Message),
	}
}

func (c *Client) Listen(ctx context.Context) {
	go c.handleMessage(ctx)
	c.msgCh <- Disconnected
}

func (c *Client) Close() error {
	close(c.done)

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		c.conn.Close(websocket.StatusNormalClosure, "")
	}

	return nil
}

func (c *Client) dial(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	u.RawQuery = "token=" + c.token

	header := make(http.Header)
	header.Add("Origin", originURL)
	opts := &websocket.DialOptions{HTTPHeader: header}

	conn, _, err := websocket.Dial(ctx, u.String(), opts)
	if err != nil {
		return err
	}
	c.conn = conn

	return nil
}

func (c *Client) listen(ctx context.Context) {
	timeout := time.NewTimer(pingInterval)
	defer timeout.Stop()

	defer func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		if c.conn != nil {
			c.conn.CloseRead(ctx)
		}
	}()

	c.msgCh <- Connected

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.done:
			return
		case <-timeout.C:
			c.msgCh <- Disconnected
			return
		default:
			msg := &message{}
			if err := wsjson.Read(ctx, c.conn, msg); err != nil {
				continue
			}

			if msg.Type == Ping {
				log.Print(msg.Type)
				timeout.Reset(pingInterval)
			}

			c.msgCh <- msg.Type
		}
	}
}

func (c *Client) handleMessage(ctx context.Context) {
	backoff := newBackoff(1*time.Second, maxBackoff)

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.done:
			return
		case msg := <-c.msgCh:
			switch msg {
			case Disconnected:
				for {
					if err := c.dial(ctx); err == nil {
						backoff.Reset()
						go c.listen(ctx)
						break
					}

					d := backoff.Duration()
					log.Printf("failed to connect, next retry in %v", d)

					select {
					case <-ctx.Done():
						return
					case <-c.done:
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
