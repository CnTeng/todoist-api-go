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

func (c *Client) DeleteSection(ctx context.Context, args *sync.SectionDeleteArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteSections(ctx context.Context, args []*sync.SectionDeleteArgs) (*sync.SyncResponse, error) {
	cmds := sync.NewCommands(args)
	return c.executeCommands(ctx, &cmds)
}

func (c *Client) ArchiveSection(ctx context.Context, args *sync.SectionArchiveArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UnarchiveSection(ctx context.Context, args *sync.SectionUnarchiveArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}
