package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) AddProject(ctx context.Context, args *sync.ProjectAddArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) UpdateProject(ctx context.Context, args *sync.ProjectUpdateArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) MoveProject(ctx context.Context, args *sync.ProjectMoveArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) MoveProjectToWorkspace(ctx context.Context, args *sync.ProjectMoveToWorkspaceArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) MoveProjectToPersonal(ctx context.Context, args *sync.ProjectMoveToPersonalArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) LeaveProject(ctx context.Context, args *sync.ProjectLeaveArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteProject(ctx context.Context, args *sync.ProjectDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) DeleteProjects(ctx context.Context, args []*sync.ProjectDeleteArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (c *Client) ArchiveProject(ctx context.Context, args *sync.ProjectArchiveArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) UnarchiveProject(ctx context.Context, args *sync.ProjectUnarchiveArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}

func (c *Client) ReorderProject(ctx context.Context, args *sync.ProjectReorderArgs) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, args)
}
