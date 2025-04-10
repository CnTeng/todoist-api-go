package ws

import (
	"context"
	"log"
	"math/rand"
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
	mu      *sync.Mutex
	conn    *websocket.Conn
	done    chan struct{}
	ch      chan struct{}
}

func NewClient(token string, handler Handler) *Client {
	return &Client{
		token:   token,
		handler: handler,
		mu:      &sync.Mutex{},
		done:    make(chan struct{}),
		ch:      make(chan struct{}, 1),
	}
}

func (c *Client) Listen(ctx context.Context) error {
	if err := c.dial(ctx); err != nil {
		return err
	}

	go c.listen(ctx)
	go c.reconnect(ctx)

	return nil
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
		select {
		case c.ch <- struct{}{}:
		default:
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.done:
			return
		case <-timeout.C:
			log.Println("connect timeout")
			return
		default:
			msg := &message{}
			if err := wsjson.Read(ctx, c.conn, msg); err != nil {
				continue
			}

			if msg.Type == ping {
				log.Println(msg.Type)
				timeout.Reset(pingInterval)
				continue
			}

			if err := c.handler.HandleMessage(ctx, msg.Type); err != nil {
				return
			}
		}
	}
}

func (c *Client) reconnect(ctx context.Context) {
	backoff := 1 * time.Second

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.done:
			return
		case <-c.ch:
			if err := c.dial(ctx); err != nil {
				log.Println("failed to connect")
				jitter := time.Duration(rand.Int63n(int64(backoff)))
				backoff = min(backoff*2+jitter, maxBackoff)

				select {
				case <-ctx.Done():
					return
				case <-c.done:
					return
				case <-time.After(backoff):
					select {
					case c.ch <- struct{}{}:
					default:
					}
				}

				continue
			}

			backoff = 1 * time.Second
			go c.listen(ctx)
		}
	}
}
