package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

type Handler interface {
	SyncToken(ctx context.Context) (*string, error)
	ResourceTypes(ctx context.Context) (*sync.ResourceTypes, error)
	HandleResponse(ctx context.Context, resp any) error
}

var DefaultHandler = &defaultHandler{syncToken: "*"}

type defaultHandler struct {
	syncToken string
}

func (h *defaultHandler) SyncToken(ctx context.Context) (*string, error) {
	return &h.syncToken, nil
}

func (h *defaultHandler) ResourceTypes(ctx context.Context) (*sync.ResourceTypes, error) {
	return &sync.ResourceTypes{sync.All}, nil
}

func (h *defaultHandler) HandleResponse(ctx context.Context, resp any) error {
	switch r := resp.(type) {
	case *sync.SyncResponse:
		h.syncToken = r.SyncToken
	}
	return nil
}
