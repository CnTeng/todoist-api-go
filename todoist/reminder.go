package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddReminder(ctx context.Context, args *sync.ReminderAddArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) UpdateReminder(ctx context.Context, args *sync.ReminderUpdateArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteReminder(ctx context.Context, args *sync.ReminderDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteReminders(ctx context.Context, args []*sync.ReminderDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommands(ctx, sync.NewCommands(args))
}
