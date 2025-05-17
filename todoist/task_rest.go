package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/rest"
	"github.com/CnTeng/todoist-api-go/sync"
)

func (s *TaskService) QuickAddTask(ctx context.Context, request *rest.TaskQuickAddRequest) (*sync.Task, error) {
	return post[sync.Task](ctx, s.client, TaskQuickAddEndpoint, request)
}

func (s *TaskService) GetCompletedTasksByCompletionDate(ctx context.Context, params *rest.TaskGetCompletedByCompletionDateParams) (*rest.TaskGetCompletedResponse, error) {
	return get[rest.TaskGetCompletedResponse](ctx, s.client, TaskCompletedByCompletionDateEndpoint, params)
}

func (s *TaskService) GetAllCompletedTasksByCompletionDate(ctx context.Context, params *rest.TaskGetCompletedByCompletionDateParams) ([]*sync.Task, error) {
	tasks := []*sync.Task{}
	for {
		resp, err := s.GetCompletedTasksByCompletionDate(ctx, params)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, resp.Tasks...)
		if resp.NextCursor == nil {
			break
		}
		params.Cursor = resp.NextCursor
	}
	return tasks, nil
}

func (s *TaskService) GetCompletedTasksByDueDate(ctx context.Context, params *rest.TaskGetCompletedByDueDateParams) (*rest.TaskGetCompletedResponse, error) {
	return get[rest.TaskGetCompletedResponse](ctx, s.client, TaskCompletedByDueDateEndpoint, params)
}

func (s *TaskService) GetAllCompletedTasksByDueDate(ctx context.Context, params *rest.TaskGetCompletedByDueDateParams) ([]*sync.Task, error) {
	tasks := []*sync.Task{}
	for {
		resp, err := s.GetCompletedTasksByDueDate(ctx, params)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, resp.Tasks...)
		if resp.NextCursor == nil {
			break
		}
		params.Cursor = resp.NextCursor
	}
	return tasks, nil
}
