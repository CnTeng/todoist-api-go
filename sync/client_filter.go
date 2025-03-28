package sync

import "context"

func (c *Client) AddFilter(ctx context.Context, args *FilterAddArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateFilter(ctx context.Context, args *FilterUpdateArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteFilter(ctx context.Context, id string) (*SyncResponse, error) {
	args := &FilterDeleteArgs{ID: id}
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderFilters(ctx context.Context, args *FilterReorderArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}
