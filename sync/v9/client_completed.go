package sync

import (
	"github.com/CnTeng/todoist-api-go/internal/utils"
	"github.com/google/go-querystring/query"
)

const getCompletedEndpoint = baseURL + "/completed/get_all"

func (sc *SyncClient) GetCompletedInfo(params *CompletedGetParams) (*CompletedGetResponse, error) {
	p, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[CompletedGetResponse](sc.client, getCompletedEndpoint, sc.token)
	return r.WithParameters(p).Post()
}
