package sync

import "context"

func (c *Client) AddLabel(ctx context.Context, args *LabelAddArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateLabel(ctx context.Context, args *LabelUpdateArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteLabel(ctx context.Context, id string) (*SyncResponse, error) {
	args := &LabelDeleteArgs{ID: id}
	return c.executeCommand(ctx, args)
}

func (c *Client) RenameSharedLabel(ctx context.Context, args *LabelRenameSharedArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteSharedLabel(ctx context.Context, name string) (*SyncResponse, error) {
	args := &LabelDeleteSharedArgs{Name: name}
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderLabels(ctx context.Context, args *LabelReorderArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}
