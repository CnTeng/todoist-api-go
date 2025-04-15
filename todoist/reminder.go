package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddReminder(ctx context.Context, args *sync.ReminderAddArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateReminder(ctx context.Context, args *sync.ReminderUpdateArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteReminder(ctx context.Context, id string) (*sync.SyncResponse, error) {
	args := &sync.ReminderDeleteArgs{ID: id}
	return c.executeCommand(ctx, args)
}
