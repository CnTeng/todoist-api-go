// Package ws provides a WebSocket client for receiving real-time sync
// notifications
package ws

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

const originURL = "https://app.todoist.com"

const (
	pingInterval = 60 * time.Second
	maxBackoff   = 128 * time.Second
)

// Client represents a Todoist WebSocket client.
type Client struct {
	url     string
	handler Handler
	logger  *log.Logger

	mu   *sync.Mutex
	conn *websocket.Conn

	wg     *sync.WaitGroup
	once   sync.Once
	cancel context.CancelFunc
	msgCh  chan Message
}

func NewClient(url string, handler Handler, logger *log.Logger) *Client {
	return &Client{
		url:     url,
		handler: handler,
		logger:  logger,
		mu:      &sync.Mutex{},
		wg:      &sync.WaitGroup{},
		once:    sync.Once{},
		msgCh:   make(chan Message),
	}
}

// Listen starts the WebSocket client to listen for incoming messages.
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

// Close stops the WebSocket client and closes the connection.
func (c *Client) Close() error {
	if c.cancel != nil {
		c.cancel()
	}

	c.wg.Wait()
	c.closeConn()

	return nil
}

// dial dials the WebSocket server and returns the connection.
func (c *Client) dial(ctx context.Context) (*websocket.Conn, error) {
	header := make(http.Header)
	header.Add("Origin", originURL)
	opts := &websocket.DialOptions{HTTPHeader: header}

	conn, _, err := websocket.Dial(ctx, c.url, opts)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// listen listens for incoming WebSocket messages.
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

			c.logger.Printf("received message: %s", msg.Type)

			if msg.Type == Ping {
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

// handleMessage handles incoming WebSocket messages and reconnection logic.
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
					c.logger.Printf("failed to connect, next retry in %s", d)

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
					c.logger.Printf("failed to handle message %s: %s", msg, err)
				}
			}
		}
	}
}

// closeConn closes the WebSocket connection.
func (c *Client) closeConn() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		c.conn.Close(websocket.StatusNormalClosure, "")
		c.conn = nil
	}
}
