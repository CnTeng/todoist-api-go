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

func (c *Client) DeleteProject(ctx context.Context, id string) (*SyncResponse, error) {
	args := &ProjectDeleteArgs{ID: id}
	return c.executeCommand(ctx, args)
}

func (c *Client) ArchiveProject(ctx context.Context, id string) (*SyncResponse, error) {
	args := &ProjectArchiveArgs{ID: id}
	return c.executeCommand(ctx, args)
}

func (c *Client) UnarchiveProject(ctx context.Context, id string) (*SyncResponse, error) {
	args := &ProjectUnarchiveArgs{ID: id}
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderProject(ctx context.Context, args *ProjectReorderArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) GetProjectInfo(ctx context.Context, params *ProjectGetInfoParams) (*ProjectGetInfoResponse, error) {
	return do[ProjectGetInfoParams, ProjectGetInfoResponse](ctx, c, projectsGetEndpoint, params)
}

func (c *Client) GetProjectData(ctx context.Context, id string) (*ProjectGetDataResponse, error) {
	return do[ProjectGetDataParams, ProjectGetDataResponse](ctx, c, projectsGetDataEndpoint, &ProjectGetDataParams{ProjectID: id})
}

func (c *Client) GetArchivedProjects(ctx context.Context, params *ProjectGetArchivedParams) (*ProjectGetArchivedResponse, error) {
	return do[ProjectGetArchivedParams, ProjectGetArchivedResponse](ctx, c, projectsGetArchivedEndpoint, params)
}
