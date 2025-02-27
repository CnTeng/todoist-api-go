package sync

import (
	"net/http"

	"github.com/CnTeng/todoist-api-go/internal/utils"
	"github.com/google/go-querystring/query"
)

const (
	baseURL     = "https://api.todoist.com"
	syncVersion = "v9"

	syncEndpoint = baseURL + "/sync/" + syncVersion + "/sync"
)

type SyncClient struct {
	token  string
	client *http.Client
}

func NewSyncClient(token string) *SyncClient {
	return &SyncClient{
		token:  token,
		client: &http.Client{},
	}
}

func (sc *SyncClient) Read(rp *ReadParams) (*Response, error) {
	p, err := query.Values(rp)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[Response](sc.client, syncEndpoint, sc.token)
	return r.WithParameters(p).Post()
}

func (sc *SyncClient) Write(wp *WriteParams) (*WriteResponse, error) {
	p, err := query.Values(wp)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[WriteResponse](sc.client, syncEndpoint, sc.token)
	return r.WithParameters(p).Post()
}
