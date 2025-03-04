package sync

import (
	"github.com/CnTeng/todoist-api-go/internal/utils"
	"github.com/google/go-querystring/query"
)

const (
	projectsGetEndpoint         = baseURL + "/projects/get"
	projectsGetDataEndpoint     = baseURL + "/projects/get_data"
	projectsGetArchivedEndpoint = baseURL + "/projects/get_archived"
)

func (sc *SyncClient) AddProject(args *ProjectAddArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) UpdateProject(args *ProjectUpdateArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) MoveProject(args *ProjectMoveArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) DeleteProject(id string) (*Response, error) {
	args := &ProjectDeleteArgs{ID: id}
	return sc.executeCommand(args)
}

func (sc *SyncClient) ArchiveProject(id string) (*Response, error) {
	args := &ProjectArchiveArgs{ID: id}
	return sc.executeCommand(args)
}

func (sc *SyncClient) UnarchiveProject(id string) (*Response, error) {
	args := &ProjectUnarchiveArgs{ID: id}
	return sc.executeCommand(args)
}

func (sc *SyncClient) ReorderProject(args *ProjectReorderArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) GetProjectInfo(params *ProjectGetInfoParams) (*ProjectGetInfoResponse, error) {
	p, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[ProjectGetInfoResponse](sc.client, projectsGetEndpoint, sc.token)
	return r.WithParameters(p).Post()
}

func (sc *SyncClient) GetProjectData(id string) (*ProjectGetDataResponse, error) {
	params := &ProjectGetDataParams{ProjectID: id}
	p, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[ProjectGetDataResponse](sc.client, projectsGetDataEndpoint, sc.token)
	return r.WithParameters(p).Post()
}

func (sc *SyncClient) GetArchivedProjects(params *ProjectGetArchivedParams) (*ProjectGetArchivedResponse, error) {
	p, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[ProjectGetArchivedResponse](sc.client, projectsGetArchivedEndpoint, sc.token)
	return r.WithParameters(p).Post()
}
