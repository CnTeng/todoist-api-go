package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/rest"
	"github.com/CnTeng/todoist-api-go/sync"
)

// Get archived projects.
func (s *ProjectService) GetArchivedProjects(ctx context.Context, params *rest.ProjectGetArchivedParams) (*rest.ProjectGetArchivedResponse, error) {
	return get[rest.ProjectGetArchivedResponse](ctx, s.client, ProjectArchivedEndpoint, params)
}

// A simple wrapper around [ProjectService.GetAllArchivedProjects] that gets all
// archived projects.
func (s *ProjectService) GetAllArchivedProjects(ctx context.Context, params *rest.ProjectGetArchivedParams) ([]*sync.Project, error) {
	ps := []*sync.Project{}
	for {
		resp, err := s.GetArchivedProjects(ctx, params)
		if err != nil {
			return nil, err
		}

		ps = append(ps, resp.Projcets...)
		if resp.NextCursor == nil {
			break
		}
		params.Cursor = resp.NextCursor
	}
	return ps, nil
}
