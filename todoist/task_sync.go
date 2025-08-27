package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

// TaskService provides methods for managing tasks.
type TaskService struct {
	client *Client
}

func NewTaskService(client *Client) *TaskService {
	return &TaskService{client: client}
}

// AddTask adds a new task to a project.
//
// See [Add a task] for more details.
//
// [Add a task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Add-a-task
func (s *TaskService) AddTask(ctx context.Context, args *sync.TaskAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// UpdateTask updates task attributes. Please note that updating the parent,
// moving, completing or uncompleting tasks is not supported by item_update,
// more specific commands have to be used instead.
//
// See [Update a task] for more details.
//
// [Update a task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Update-a-task
func (s *TaskService) UpdateTask(ctx context.Context, args *sync.TaskUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// UpdateTaskDayOrders updates the day orders of multiple tasks at once.
//
// See [Update day orders] for more details.
//
// [Update day orders]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Update-day-orders
func (s *TaskService) UpdateTaskDayOrders(ctx context.Context, args *sync.TaskUpdateDayOrdersArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// MoveTask moves task to a different location. Only one of parent_id,
// section_id or project_id must be set.
//
// See [Move a task] for more details.
//
// [Move a task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Move-a-task
func (s *TaskService) MoveTask(ctx context.Context, args *sync.TaskMoveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// ReorderTasks updates child_order properties of items in bulk.
//
// See [Reorder tasks] for more details.
//
// [Reorder tasks]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Reorder-tasks
func (s *TaskService) ReorderTasks(ctx context.Context, args *sync.TaskReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// CompleteTask completes a task and its sub-tasks and moves them to the
// archive. See also [TaskCloseArgs] for a simplified version of the command.
//
// See [Complete task] for more details.
//
// [Complete task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Complete-task
func (s *TaskService) CompleteTask(ctx context.Context, args *sync.TaskCompleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// CompleteTaskRecurring completes a recurring task. The reason why this is a
// special case is because we need to mark a recurring completion (and using
// item_update won't do this). See also [TaskCloseArgs] for a simplified version
// of the command.
//
// See [Complete a recurring task] for more details.
//
// [Complete a recurring task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Complete-a-recurring-task
func (s *TaskService) CompleteTaskRecurring(ctx context.Context, args *sync.TaskCompleteRecurringArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// CloseTask is a simplified version of item_complete /
// item_update_date_complete. The command does exactly what official clients do
// when you close a task:
//
//   - regular tasks are completed and moved to the archive
//   - recurring tasks are scheduled to their next occurrence
//
// See [Close task] for more details.
//
// [Close task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Close-task
func (s *TaskService) CloseTask(ctx context.Context, args *sync.TaskCloseArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// CloseTasks is a simple wrapper around [TaskService.CloseTask] that closes multiple tasks.
func (s *TaskService) CloseTasks(ctx context.Context, args []*sync.TaskCloseArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

// UncompleteTask uncompletes and restores an archived item.
//
// Any ancestor items or sections will also be reinstated. Items will have the
// checked value reset.
//
// The reinstated items and sections will appear at the end of the list within
// their parent, after any previously active tasks.
//
// See [Uncomplete task] for more details.
//
// [Uncomplete task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Uncomplete-item
func (s *TaskService) UncompleteTask(ctx context.Context, args *sync.TaskUncompleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// UncompleteTasks is a simple wrapper around [TaskService.UncompleteTask] that
// uncompletes multiple tasks.
func (s *TaskService) UncompleteTasks(ctx context.Context, args []*sync.TaskUncompleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

// DeleteTask deletes a task and all its sub-tasks.
//
// See [Delete a task] for more details.
//
// [Delete a task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Delete-tasks
func (s *TaskService) DeleteTask(ctx context.Context, args *sync.TaskDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// DeleteTasks is a simple wrapper around [TaskService.DeleteTask] that deletes
// multiple.
func (s *TaskService) DeleteTasks(ctx context.Context, args []*sync.TaskDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
