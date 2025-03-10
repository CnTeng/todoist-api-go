package sync

import (
	"context"
	"net/http"

	"github.com/CnTeng/todoist-api-go/internal/utils"
	"github.com/google/go-querystring/query"
)

const (
	baseURL = "https://api.todoist.com/sync/v9"

	syncEndpoint = baseURL + "/sync"
)

type Handler interface {
	SyncToken() (*string, error)
	ResourceTypes() (*ResourceTypes, error)
	HandleResponse(resp any) error
}

var DefaultHandler = &defaultHandler{syncToken: "*"}

type defaultHandler struct {
	syncToken string
}

func (h *defaultHandler) SyncToken() (*string, error) {
	return &h.syncToken, nil
}

func (h *defaultHandler) ResourceTypes() (*ResourceTypes, error) {
	return &ResourceTypes{All}, nil
}

func (h *defaultHandler) HandleResponse(resp any) error {
	switch r := resp.(type) {
	case *SyncResponse:
		h.syncToken = r.SyncToken
	}
	return nil
}

type Client struct {
	client  *http.Client
	token   string
	handler Handler
}

func NewClient(client *http.Client, token string) *Client {
	return NewClientWithHandler(client, token, DefaultHandler)
}

func NewClientWithHandler(client *http.Client, token string, handler Handler) *Client {
	return &Client{
		client:  client,
		token:   token,
		handler: handler,
	}
}

func do[T, U any](ctx context.Context, c *Client, endpoint string, params *T) (*U, error) {
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[U](c.client, endpoint, c.token)
	resp, err := r.WithParameters(v).Post(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.handler.HandleResponse(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) do(ctx context.Context, p *SyncParams) (*SyncResponse, error) {
	return do[SyncParams, SyncResponse](ctx, c, syncEndpoint, p)
}

func (c *Client) executeCommands(ctx context.Context, cmds *Commands) (*SyncResponse, error) {
	st, err := c.handler.SyncToken()
	if err != nil {
		return nil, err
	}

	rt, err := c.handler.ResourceTypes()
	if err != nil {
		return nil, err
	}

	p := &SyncParams{SyncToken: st, ResourceTypes: rt, Commands: cmds}
	return c.do(ctx, p)
}

func (c *Client) executeCommand(ctx context.Context, args CommandArgs) (*SyncResponse, error) {
	cmd := NewCommand(args)
	return c.executeCommands(ctx, &Commands{cmd})
}

func (c *Client) Do(ctx context.Context, p *SyncParams) (*SyncResponse, error) {
	return c.do(ctx, p)
}

func (c *Client) Sync(ctx context.Context) (*SyncResponse, error) {
	st, err := c.handler.SyncToken()
	if err != nil {
		return nil, err
	}

	rt, err := c.handler.ResourceTypes()
	if err != nil {
		return nil, err
	}

	p := &SyncParams{SyncToken: st, ResourceTypes: rt}
	return c.do(ctx, p)
}

func (c *Client) SyncForced(ctx context.Context) (*SyncResponse, error) {
	st := "*"

	rt, err := c.handler.ResourceTypes()
	if err != nil {
		return nil, err
	}

	p := &SyncParams{SyncToken: &st, ResourceTypes: rt}
	return c.do(ctx, p)
}

func (c *Client) ExecuteCommand(ctx context.Context, args CommandArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) ExecuteCommands(ctx context.Context, cmds *Commands) (*SyncResponse, error) {
	return c.executeCommands(ctx, cmds)
}
