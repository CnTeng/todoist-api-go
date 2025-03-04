package sync

import (
	"net/http"

	"github.com/CnTeng/todoist-api-go/internal/utils"
	"github.com/google/go-querystring/query"
)

const (
	baseURL = "https://api.todoist.com/sync/v9"

	syncEndpoint = baseURL + "/sync"
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

func (sc *SyncClient) Write(cs *Commands) (*Response, error) {
	wp := &WriteParams{Commands: cs}
	p, err := query.Values(wp)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[Response](sc.client, syncEndpoint, sc.token)
	return r.WithParameters(p).Post()
}

func (sc *SyncClient) executeCommand(args CommandArgs) (*Response, error) {
	c := NewCommand(args)

	resp, err := sc.Write(&Commands{c})
	if err != nil {
		return nil, err
	}

	err = resp.SyncStatus[c.UUID]
	if err != nil {
		return nil, err
	}

	return resp, nil
}
