package sync

import "context"

const (
	projectsGetEndpoint         = baseURL + "/projects/get"
	projectsGetDataEndpoint     = baseURL + "/projects/get_data"
	projectsGetArchivedEndpoint = baseURL + "/projects/get_archived"
)

func (c *Client) AddProject(ctx context.Context, args *ProjectAddArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateProject(ctx context.Context, args *ProjectUpdateArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) MoveProject(ctx context.Context, args *ProjectMoveArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteProject(ctx context.Context, args *ProjectDeleteArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteProjects(ctx context.Context, args []*ProjectDeleteArgs) (*SyncResponse, error) {
	cmds := make(Commands, 0, len(args))
	for _, arg := range args {
		cmds = append(cmds, NewCommand(arg))
	}
	return c.executeCommands(ctx, &cmds)
}

func (c *Client) ArchiveProject(ctx context.Context, args *ProjectArchiveArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UnarchiveProject(ctx context.Context, args *ProjectUnarchiveArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderProject(ctx context.Context, args *ProjectReorderArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) GetProjectInfo(ctx context.Context, params *ProjectGetInfoParams) (*ProjectGetInfoResponse, error) {
	return do[ProjectGetInfoParams, ProjectGetInfoResponse](ctx, c, projectsGetEndpoint, params)
}

func (c *Client) GetProjectData(ctx context.Context, params *ProjectGetDataParams) (*ProjectGetDataResponse, error) {
	return do[ProjectGetDataParams, ProjectGetDataResponse](ctx, c, projectsGetDataEndpoint, params)
}

func (c *Client) GetArchivedProjects(ctx context.Context, params *ProjectGetArchivedParams) (*ProjectGetArchivedResponse, error) {
	return do[ProjectGetArchivedParams, ProjectGetArchivedResponse](ctx, c, projectsGetArchivedEndpoint, params)
}
