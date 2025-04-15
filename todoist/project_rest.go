package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/rest"
)

func (c *Client) GetProjectsArchived(ctx context.Context, params *rest.ProjectGetArchivedParams) (*rest.ProjectGetArchivedResponse, error) {
	return get[rest.ProjectGetArchivedResponse](ctx, c, ProjectArchivedEndpoint, params)
}
