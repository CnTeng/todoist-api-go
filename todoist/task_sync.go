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

func (s *TaskService) AddTask(ctx context.Context, args *sync.TaskAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) UpdateTask(ctx context.Context, args *sync.TaskUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) UpdateTaskDayOrders(ctx context.Context, args *sync.TaskUpdateDayOrdersArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) MoveTask(ctx context.Context, args *sync.TaskMoveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) ReorderTasks(ctx context.Context, args *sync.TaskReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) CompleteTask(ctx context.Context, args *sync.TaskCompleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) CompleteTaskRecurring(ctx context.Context, args *sync.TaskCompleteRecurringArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) CloseTask(ctx context.Context, args *sync.TaskCloseArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) CloseTasks(ctx context.Context, args []*sync.TaskCloseArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (s *TaskService) UncompleteTask(ctx context.Context, args *sync.TaskUncompleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) UncompleteTasks(ctx context.Context, args []*sync.TaskUncompleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (s *TaskService) DeleteTask(ctx context.Context, args *sync.TaskDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *TaskService) DeleteTasks(ctx context.Context, args []*sync.TaskDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
