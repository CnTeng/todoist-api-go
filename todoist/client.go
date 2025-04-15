package todoist

import (
	"context"
	"net/http"

	"github.com/CnTeng/todoist-api-go/sync"
	"github.com/google/go-querystring/query"
)

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

func get[T any](ctx context.Context, c *Client, endpoint string, params any) (*T, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	r := NewRequest[T](c.client, endpoint, c.token)
	resp, err := r.WithParameters(v).Get(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.handler.HandleResponse(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func post[T any](ctx context.Context, c *Client, endpoint string, body any) (*T, error) {
	r := NewRequest[T](c.client, endpoint, c.token)
	resp, err := r.WithBody(body).Post(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.handler.HandleResponse(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func do[T, U any](ctx context.Context, c *Client, endpoint string, params *T) (*U, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	r := NewRequest[U](c.client, endpoint, c.token)
	resp, err := r.WithParameters(v).Post(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.handler.HandleResponse(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) do(ctx context.Context, p *sync.SyncParams) (*sync.SyncResponse, error) {
	return do[sync.SyncParams, sync.SyncResponse](ctx, c, syncEndpoint, p)
}

func (c *Client) executeCommands(ctx context.Context, cmds *sync.Commands) (*sync.SyncResponse, error) {
	st, err := c.handler.SyncToken()
	if err != nil {
		return nil, err
	}

	rt, err := c.handler.ResourceTypes()
	if err != nil {
		return nil, err
	}

	p := &sync.SyncParams{SyncToken: st, ResourceTypes: rt, Commands: cmds}
	return c.do(ctx, p)
}

func (c *Client) executeCommand(ctx context.Context, args sync.CommandArgs) (*sync.SyncResponse, error) {
	cmd := sync.NewCommand(args)
	return c.executeCommands(ctx, &sync.Commands{cmd})
}

func (c *Client) Do(ctx context.Context, p *sync.SyncParams) (*sync.SyncResponse, error) {
	return c.do(ctx, p)
}

func (c *Client) Sync(ctx context.Context, isForce bool) (*sync.SyncResponse, error) {
	var syncToken string

	if isForce {
		syncToken = sync.DefaultSyncToken
	} else {
		st, err := c.handler.SyncToken()
		if err != nil {
			return nil, err
		}
		syncToken = *st
	}

	rt, err := c.handler.ResourceTypes()
	if err != nil {
		return nil, err
	}

	p := &sync.SyncParams{SyncToken: &syncToken, ResourceTypes: rt}
	return c.do(ctx, p)
}

func (c *Client) ExecuteCommand(ctx context.Context, args sync.CommandArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) ExecuteCommands(ctx context.Context, cmds *sync.Commands) (*sync.SyncResponse, error) {
	return c.executeCommands(ctx, cmds)
}
