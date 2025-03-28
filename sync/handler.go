package sync

type Handler interface {
	SyncToken() (*string, error)
	ResourceTypes() (*ResourceTypes, error)
	HandleResponse(resp any) error
}

var DefaultHandler = &defaultHandler{syncToken: "*"}

type defaultHandler struct {
	syncToken string
}

func (h *defaultHandler) SyncToken() (*string, error) {
	return &h.syncToken, nil
}

func (h *defaultHandler) ResourceTypes() (*ResourceTypes, error) {
	return &ResourceTypes{All}, nil
}

func (h *defaultHandler) HandleResponse(resp any) error {
	switch r := resp.(type) {
	case *SyncResponse:
		h.syncToken = r.SyncToken
	}
	return nil
}
