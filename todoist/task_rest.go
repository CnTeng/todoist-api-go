package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/rest"
	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddTaskQuick(ctx context.Context, request *rest.TaskQuickAddRequest) (*sync.Task, error) {
	return post[sync.Task](ctx, c, TaskQuickAddEndpoint, request)
}

func (c *Client) GetTasksCompletedByCompletionDate(ctx context.Context, params *rest.TaskGetCompletedByCompletionDateParams) (*rest.TaskGetCompletedResponse, error) {
	return get[rest.TaskGetCompletedResponse](ctx, c, TaskCompletedByCompletionDateEndpoint, params)
}

func (c *Client) GetTasksCompletedByDueDate(ctx context.Context, params *rest.TaskGetCompletedByDueDateParams) (*rest.TaskGetCompletedResponse, error) {
	return get[rest.TaskGetCompletedResponse](ctx, c, TaskCompletedByDueDateEndpoint, params)
}
