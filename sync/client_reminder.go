package sync

import "context"

func (c *Client) AddReminder(ctx context.Context, args *ReminderAddArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateReminder(ctx context.Context, args *ReminderUpdateArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteReminder(ctx context.Context, id string) (*SyncResponse, error) {
	args := &ReminderDeleteArgs{ID: id}
	return c.executeCommand(ctx, args)
}
