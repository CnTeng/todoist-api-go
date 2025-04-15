package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddSection(ctx context.Context, args *sync.SectionAddArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateSection(ctx context.Context, args *sync.SectionUpdateArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) MoveSection(ctx context.Context, args *sync.SectionMoveArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderSection(ctx context.Context, args *sync.SectionReorderArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteSection(ctx context.Context, id string) (*sync.SyncResponse, error) {
	args := &sync.SectionDeleteArgs{id}
	return c.executeCommand(ctx, args)
}

func (c *Client) ArchiveSection(ctx context.Context, id string) (*sync.SyncResponse, error) {
	args := &sync.SectionArchiveArgs{ID: id}
	return c.executeCommand(ctx, args)
}

func (c *Client) UnarchiveSection(ctx context.Context, id string) (*sync.SyncResponse, error) {
	args := &sync.SectionUnarchiveArgs{ID: id}
	return c.executeCommand(ctx, args)
}
