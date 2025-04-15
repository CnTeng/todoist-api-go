package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddFilter(ctx context.Context, args *sync.FilterAddArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) UpdateFilter(ctx context.Context, args *sync.FilterUpdateArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteFilter(ctx context.Context, args *sync.FilterDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteFilters(ctx context.Context, args []*sync.FilterDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (c *Client) ReorderFilters(ctx context.Context, args *sync.FilterReorderArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}
