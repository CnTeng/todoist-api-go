package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

const (
	projectsGetEndpoint         = baseURL + "/projects/get"
	projectsGetDataEndpoint     = baseURL + "/projects/get_data"
	projectsGetArchivedEndpoint = baseURL + "/projects/get_archived"
)

func (c *Client) AddProject(ctx context.Context, args *sync.ProjectAddArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateProject(ctx context.Context, args *sync.ProjectUpdateArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) MoveProject(ctx context.Context, args *sync.ProjectMoveArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteProject(ctx context.Context, args *sync.ProjectDeleteArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteProjects(ctx context.Context, args []*sync.ProjectDeleteArgs) (*sync.SyncResponse, error) {
	cmds := make(sync.Commands, 0, len(args))
	for _, arg := range args {
		cmds = append(cmds, sync.NewCommand(arg))
	}
	return c.executeCommands(ctx, &cmds)
}

func (c *Client) ArchiveProject(ctx context.Context, args *sync.ProjectArchiveArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UnarchiveProject(ctx context.Context, args *sync.ProjectUnarchiveArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderProject(ctx context.Context, args *sync.ProjectReorderArgs) (*sync.SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) GetProjectInfo(ctx context.Context, params *sync.ProjectGetInfoParams) (*sync.ProjectGetInfoResponse, error) {
	return do[sync.ProjectGetInfoParams, sync.ProjectGetInfoResponse](ctx, c, projectsGetEndpoint, params)
}

func (c *Client) GetProjectData(ctx context.Context, params *sync.ProjectGetDataParams) (*sync.ProjectGetDataResponse, error) {
	return do[sync.ProjectGetDataParams, sync.ProjectGetDataResponse](ctx, c, projectsGetDataEndpoint, params)
}

func (c *Client) GetArchivedProjects(ctx context.Context, params *sync.ProjectGetArchivedParams) (*sync.ProjectGetArchivedResponse, error) {
	return do[sync.ProjectGetArchivedParams, sync.ProjectGetArchivedResponse](ctx, c, projectsGetArchivedEndpoint, params)
}
