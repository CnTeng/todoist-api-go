package sync

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/google/uuid"
)

const syncStatusOk = "ok"

type WriteParams struct {
	Commands *Commands `json:"commands" url:"commands"`
}

type WriteResponse struct {
	SyncStatus    map[uuid.UUID]error  `json:"sync_status"`
	TempIDMapping map[uuid.UUID]string `json:"temp_id_mapping"`
	FullSync      bool                 `json:"full_sync"`
	SyncToken     string               `json:"sync_token"`
}

func (wr *WriteResponse) UnmarshalJSON(data []byte) error {
	type alias WriteResponse
	aux := &struct {
		SyncStatus map[uuid.UUID]any `json:"sync_status"`
		*alias
	}{alias: (*alias)(wr)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	wr.SyncStatus = make(map[uuid.UUID]error, len(aux.SyncStatus))
	for k, v := range aux.SyncStatus {
		switch v := v.(type) {
		case string:
			if v == syncStatusOk {
				wr.SyncStatus[k] = nil
			} else {
				wr.SyncStatus[k] = errors.New(strings.ToLower(v))
			}
		case map[string]any:
			if err, ok := v["error"]; ok {
				if err, ok := err.(string); ok {
					wr.SyncStatus[k] = errors.New(strings.ToLower(err))
				}
			} else {
				wr.SyncStatus[k] = errors.New("unknown error")
			}
		default:
			wr.SyncStatus[k] = errors.New("unknown error")
		}
	}

	return nil
}
