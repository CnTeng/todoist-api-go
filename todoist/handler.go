package todoist

import "github.com/CnTeng/todoist-api-go/sync"

type Handler interface {
	SyncToken() (*string, error)
	ResourceTypes() (*sync.ResourceTypes, error)
	HandleResponse(resp any) error
}

var DefaultHandler = &defaultHandler{syncToken: "*"}

type defaultHandler struct {
	syncToken string
}

func (h *defaultHandler) SyncToken() (*string, error) {
	return &h.syncToken, nil
}

func (h *defaultHandler) ResourceTypes() (*sync.ResourceTypes, error) {
	return &sync.ResourceTypes{sync.All}, nil
}

func (h *defaultHandler) HandleResponse(resp any) error {
	switch r := resp.(type) {
	case *sync.SyncResponse:
		h.syncToken = r.SyncToken
	}
	return nil
}
