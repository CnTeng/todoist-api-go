package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddTask(ctx context.Context, args *sync.TaskAddArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) UpdateTask(ctx context.Context, args *sync.TaskUpdateArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) MoveTask(ctx context.Context, args *sync.TaskMoveArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) ReorderTasks(ctx context.Context, args *sync.TaskReorderArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteTask(ctx context.Context, args *sync.TaskDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteTasks(ctx context.Context, args []*sync.TaskDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (c *Client) CompleteTask(ctx context.Context, args *sync.TaskCompleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) CompleteTaskRecurring(ctx context.Context, args *sync.TaskCompleteRecurringArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) UncompleteTask(ctx context.Context, args *sync.TaskUncompleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) CloseTask(ctx context.Context, args *sync.TaskCloseArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) UpdateTaskDayOrders(ctx context.Context, args *sync.TaskUpdateDayOrdersArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}
