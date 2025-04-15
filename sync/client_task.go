package sync

import (
	"context"

	"github.com/CnTeng/todoist-api-go/rest"
)

const (
	taskQuickAddEndpoint              = baseURL + "/tasks/quick"
	taskGetCompletedByDueDateEndpoint = baseURL + "/tasks/completed"
)

func (c *Client) AddTask(ctx context.Context, args *TaskAddArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateTask(ctx context.Context, args *TaskUpdateArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) MoveTask(ctx context.Context, args *TaskMoveArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderTasks(ctx context.Context, args *TaskReorderArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteTask(ctx context.Context, args *TaskDeleteArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteTasks(ctx context.Context, args []*TaskDeleteArgs) (*SyncResponse, error) {
	cmds := make(Commands, 0, len(args))
	for _, arg := range args {
		cmds = append(cmds, NewCommand(arg))
	}
	return c.executeCommands(ctx, &cmds)
}

func (c *Client) CompleteTask(ctx context.Context, args *TaskCompleteArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) CompleteTaskRecurring(ctx context.Context, args *TaskCompleteRecurringArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UncompleteTask(ctx context.Context, args *TaskUncompleteArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) CloseTask(ctx context.Context, args *TaskCloseArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateTaskDayOrders(ctx context.Context, args *TaskUpdateDayOrdersArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) AddTaskQuick(ctx context.Context, request *rest.TaskQuickAddRequest) (*Task, error) {
	return post[Task](ctx, c, taskQuickAddEndpoint, request)
}

func (c *Client) GetTasksCompleted(ctx context.Context, params *TaskGetCompletedByDueDateParams) (*TaskGetCompletedResponse, error) {
	return get[TaskGetCompletedResponse](ctx, c, taskGetCompletedByDueDateEndpoint, params)
}
