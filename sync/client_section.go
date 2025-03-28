package sync

import "context"

func (c *Client) AddSection(ctx context.Context, args *SectionAddArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateSection(ctx context.Context, args *SectionUpdateArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) MoveSection(ctx context.Context, args *SectionMoveArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderSection(ctx context.Context, args *SectionReorderArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteSection(ctx context.Context, id string) (*SyncResponse, error) {
	args := &SectionDeleteArgs{id}
	return c.executeCommand(ctx, args)
}

func (c *Client) ArchiveSection(ctx context.Context, id string) (*SyncResponse, error) {
	args := &SectionArchiveArgs{ID: id}
	return c.executeCommand(ctx, args)
}

func (c *Client) UnarchiveSection(ctx context.Context, id string) (*SyncResponse, error) {
	args := &SectionUnarchiveArgs{ID: id}
	return c.executeCommand(ctx, args)
}
