package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddLabel(ctx context.Context, args *sync.LabelAddArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateLabel(ctx context.Context, args *sync.LabelUpdateArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteLabel(ctx context.Context, id string) (*sync.SyncResponse, error) {
	args := &sync.LabelDeleteArgs{ID: id}
	return c.executeCommand(ctx, args)
}

func (c *Client) RenameSharedLabel(ctx context.Context, args *sync.LabelRenameSharedArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteSharedLabel(ctx context.Context, name string) (*sync.SyncResponse, error) {
	args := &sync.LabelDeleteSharedArgs{Name: name}
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderLabels(ctx context.Context, args *sync.LabelReorderArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}
