package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddLabel(ctx context.Context, args *sync.LabelAddArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) UpdateLabel(ctx context.Context, args *sync.LabelUpdateArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteLabel(ctx context.Context, args *sync.LabelDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteLabels(ctx context.Context, args []*sync.LabelDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (c *Client) RenameLabel(ctx context.Context, args *sync.LabelRenameArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteLabelOccurrences(ctx context.Context, args *sync.LabelDeleteOccurrencesArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) ReorderLabels(ctx context.Context, args *sync.LabelReorderArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}
