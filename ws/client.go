package ws

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/coder/websocket"
)

const (
	baseURL   = "wss://ws.todoist.com/ws"
	originURL = "https://app.todoist.com"
)

type notification struct {
	Type Notification `json:"type"`
}

type Notification string

const (
	Ping            Notification = "ping"
	SyncNeeded      Notification = "sync_needed"
	CalenderUpdated Notification = "calendar_updated"
)

type Client struct {
	token   string
	handler Handler
}

func NewClient(token string, handler Handler) *Client {
	return &Client{
		token:   token,
		handler: handler,
	}
}

func (c *Client) Listen(ctx context.Context) error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	u.RawQuery = "token=" + c.token

	header := make(http.Header)
	header.Add("Origin", originURL)
	opts := &websocket.DialOptions{HTTPHeader: header}

	go func() {
		conn, _, err := websocket.Dial(ctx, u.String(), opts)
		if err != nil {
			return
		}
		defer func() { _ = conn.CloseNow() }()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, msg, err := conn.Read(ctx)
				if err != nil {
					continue
				}

				noti := &notification{}
				if err := json.Unmarshal(msg, noti); err != nil {
					return
				}

				if err := c.handler.HandleNotification(ctx, noti.Type); err != nil {
					return
				}
			}
		}
	}()

	return nil
}
