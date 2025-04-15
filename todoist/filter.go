package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddFilter(ctx context.Context, args *sync.FilterAddArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateFilter(ctx context.Context, args *sync.FilterUpdateArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteFilter(ctx context.Context, args *sync.FilterDeleteArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteFilters(ctx context.Context, args []*sync.FilterDeleteArgs) (*sync.SyncResponse, error) {
	cmds := sync.NewCommands(args)
	return c.executeCommands(ctx, &cmds)
}

func (c *Client) ReorderFilters(ctx context.Context, args *sync.FilterReorderArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}
