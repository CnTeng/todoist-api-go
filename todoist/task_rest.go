package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/rest"
	"github.com/CnTeng/todoist-api-go/sync"
)

// Add a new task using the Quick Add implementation similar to that used in the
// official clients.
//
// See [Quick add a task] for more details.
//
// [Quick add a task]: https://todoist.com/api/v1/docs#tag/Tasks/operation/quick_add_api_v1_tasks_quick_post
func (s *TaskService) QuickAddTask(ctx context.Context, request *rest.TaskQuickAddRequest) (*sync.Task, error) {
	return post[sync.Task](ctx, s.client, TaskQuickAddEndpoint, request)
}

// Retrieves a list of completed tasks strictly limited by the specified
// completion date range (up to 3 months).
//
// It can retrieve completed items:
//
//   - From all the projects the user has joined in a workspace
//   - From all the projects of the user
//   - That match many [supported filters]
//
// By default, the response is limited to a page containing a maximum of 50
// items (configurable using limit).
//
// Subsequent pages of results can be fetched by using the next_cursor value
// from the response as the cursor value for the next request.
//
// See [Tasks Completed By Completion Date] for more details.
//
// [supported filters]: https://todoist.com/help/articles/introduction-to-filters-V98wIH
// [Tasks Completed By Completion Date]: https://todoist.com/api/v1/docs#tag/Tasks/operation/tasks_completed_by_completion_date_api_v1_tasks_completed_by_completion_date_get
func (s *TaskService) GetCompletedTasksByCompletionDate(ctx context.Context, params *rest.TaskGetCompletedByCompletionDateParams) (*rest.TaskGetCompletedResponse, error) {
	return get[rest.TaskGetCompletedResponse](ctx, s.client, TaskCompletedByCompletionDateEndpoint, params)
}

// A simple wrapper around [TaskService.GetAllCompletedTasksByCompletionDate]
// that retrieves all.
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

// Retrieves a list of completed items strictly limited by the specified due
// date range (up to 6 weeks).
//
// It can retrieve completed items:
//
//   - From within a project, section, or parent item
//   - From all the projects the user has joined in a workspace
//   - From all the projects of the user
//   - That match many [supported filters]
//
// By default, the response is limited to a page containing a maximum of 50
// items (configurable using limit).
//
// Subsequent pages of results can be fetched by using the next_cursor value
// from the response as the cursor value for the next request.
//
// See [Tasks Completed By Due Date] for more details.
//
// [supported filters]: https://todoist.com/help/articles/introduction-to-filters-V98wIH
// [Tasks Completed By Due Date]: https://todoist.com/api/v1/docs#tag/Tasks/operation/tasks_completed_by_due_date_api_v1_tasks_completed_by_due_date_get
func (s *TaskService) GetCompletedTasksByDueDate(ctx context.Context, params *rest.TaskGetCompletedByDueDateParams) (*rest.TaskGetCompletedResponse, error) {
	return get[rest.TaskGetCompletedResponse](ctx, s.client, TaskCompletedByDueDateEndpoint, params)
}

// A simple wrapper around [TaskService.GetCompletedTasksByDueDate] that
// retrieves all.
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
