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
	SyncToken() *string
	ResourceTypes() *ResourceTypes
	HandleResponse(resp any) error
}

type defaultHandler struct {
	syncToken *string
}

func newDefaultHandler() *defaultHandler {
	token := "*"
	return &defaultHandler{
		syncToken: &token,
	}
}

func (h *defaultHandler) SyncToken() *string {
	return h.syncToken
}

func (h *defaultHandler) ResourceTypes() *ResourceTypes {
	return &ResourceTypes{"all"}
}

func (h *defaultHandler) HandleResponse(resp any) error {
	switch r := resp.(type) {
	case *SyncResponse:
		h.syncToken = &r.SyncToken
	}
	return nil
}

type Client struct {
	client  *http.Client
	token   string
	handler Handler
}

func NewClient(client *http.Client, token string) *Client {
	return NewClientWithHandler(client, token, newDefaultHandler())
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
	p := &SyncParams{SyncToken: c.handler.SyncToken(), ResourceTypes: c.handler.ResourceTypes(), Commands: cmds}
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
	p := &SyncParams{SyncToken: c.handler.SyncToken(), ResourceTypes: c.handler.ResourceTypes()}
	return c.do(ctx, p)
}

func (c *Client) SyncForced(ctx context.Context) (*SyncResponse, error) {
	token := "*"
	p := &SyncParams{SyncToken: &token, ResourceTypes: c.handler.ResourceTypes()}
	return c.do(ctx, p)
}

func (c *Client) ExecuteCommand(ctx context.Context, args CommandArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) ExecuteCommands(ctx context.Context, cmds *Commands) (*SyncResponse, error) {
	return c.executeCommands(ctx, cmds)
}
