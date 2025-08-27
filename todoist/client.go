// Package todoist implements a client for interacting with the Todoist API v1.
package todoist

import (
	"context"
	"net/http"

	"github.com/CnTeng/todoist-api-go/sync"
)

// Client is a client to interact with the Todoist API v1.
type Client struct {
	client  *http.Client
	token   string
	handler Handler
}

func NewClient(client *http.Client, token string, handler Handler) *Client {
	return &Client{
		client:  client,
		token:   token,
		handler: handler,
	}
}

// Sync data between the Todoist client.
func (c *Client) Sync(ctx context.Context, request *sync.SyncRequest) (*sync.SyncResponse, error) {
	return post[sync.SyncResponse](ctx, c, syncEndpoint, request)
}

// SyncWithAutoToken performs a sync operation with automatic handling of the
// sync token.
func (c *Client) SyncWithAutoToken(ctx context.Context, fullSync bool) (*sync.SyncResponse, error) {
	syncToken := sync.DefaultSyncToken
	if !fullSync {
		st, err := c.handler.SyncToken(ctx)
		if err != nil {
			return nil, err
		}
		syncToken = *st
	}

	rt, err := c.handler.ResourceTypes(ctx)
	if err != nil {
		return nil, err
	}

	req := &sync.SyncRequest{SyncToken: &syncToken, ResourceTypes: rt}
	return c.Sync(ctx, req)
}

// ExecuteCommand executes a single command in a single sync request.
func (c *Client) ExecuteCommand(ctx context.Context, args sync.CommandArgs) (*sync.SyncResponse, error) {
	cmd := sync.NewCommand(args)
	return c.ExecuteCommands(ctx, sync.Commands{cmd})
}

// ExecuteCommands executes multiple commands in a single sync request.
func (c *Client) ExecuteCommands(ctx context.Context, cmds sync.Commands) (*sync.SyncResponse, error) {
	st, err := c.handler.SyncToken(ctx)
	if err != nil {
		return nil, err
	}

	rt, err := c.handler.ResourceTypes(ctx)
	if err != nil {
		return nil, err
	}

	req := &sync.SyncRequest{SyncToken: st, ResourceTypes: rt, Commands: cmds}
	return c.Sync(ctx, req)
}

func get[T any](ctx context.Context, c *Client, endpoint string, params any) (*T, error) {
	r := newRequest[T](c.client, endpoint, c.token)
	resp, err := r.withParams(params).do(ctx, http.MethodGet)
	if err != nil {
		return nil, err
	}

	if err := c.handler.HandleResponse(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func post[T any](ctx context.Context, c *Client, endpoint string, body any) (*T, error) {
	r := newRequest[T](c.client, endpoint, c.token)
	resp, err := r.withBody(body).do(ctx, http.MethodPost)
	if err != nil {
		return nil, err
	}

	if err := c.handler.HandleResponse(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
