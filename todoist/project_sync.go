package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

// ProjectSync provides methods to manage projects in Todoist.
type ProjectService struct {
	client *Client
}

func NewProjectService(client *Client) *ProjectService {
	return &ProjectService{client: client}
}

func (s *ProjectService) AddProject(ctx context.Context, args *sync.ProjectAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) UpdateProject(ctx context.Context, args *sync.ProjectUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) MoveProject(ctx context.Context, args *sync.ProjectMoveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) MoveProjectToWorkspace(ctx context.Context, args *sync.ProjectMoveToWorkspaceArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) MoveProjectToPersonal(ctx context.Context, args *sync.ProjectMoveToPersonalArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) LeaveProject(ctx context.Context, args *sync.ProjectLeaveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) ReorderProjects(ctx context.Context, args *sync.ProjectReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) ArchiveProject(ctx context.Context, args *sync.ProjectArchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) ArchiveProjects(ctx context.Context, args []*sync.ProjectArchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (s *ProjectService) UnarchiveProject(ctx context.Context, args *sync.ProjectUnarchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) UnarchiveProjects(ctx context.Context, args []*sync.ProjectUnarchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (s *ProjectService) DeleteProject(ctx context.Context, args *sync.ProjectDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *ProjectService) DeleteProjects(ctx context.Context, args []*sync.ProjectDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
