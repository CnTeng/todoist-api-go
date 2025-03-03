package sync

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/google/uuid"
)

const syncStatusOk = "ok"

type Response struct {
	// TODO:
	CompletedInfo []any `json:"completed_info"`

	// TODO:
	Collaborators []any `json:"collaborators"`

	// TODO:
	CollaboratorStates []any `json:"collaborator_states"`

	// TODO:
	DayOrders map[string]any `json:"day_orders"`

	Filters []*Filter `json:"filters"`

	FullSync bool `json:"full_sync"`

	Items []*Item `json:"items"`

	ItemNotes []*ItemNote `json:"notes"`

	Labels []*Label `json:"labels"`

	// TODO:
	LiveNotifications []any `json:"live_notifications"`

	// TODO:
	LiveNotificationsLastReadID string `json:"live_notifications_last_read_id"`

	// TODO:
	Locations []any `json:"locations"`

	ProjectNotes []*ProjectNote `json:"project_notes"`

	Projects []*Project `json:"projects"`

	Reminders []*Reminder `json:"reminders"`

	// TODO:
	Sections []*Section `json:"sections"`

	// TODO:
	Stats map[string]any `json:"stats"`

	// TODO:
	SettingsNotifications map[string]any `json:"settings_notifications"`

	SyncStatus map[uuid.UUID]error `json:"sync_status"`

	SyncToken string `json:"sync_token"`

	// TODO:
	TempIDMapping map[uuid.UUID]string `json:"temp_id_mapping"`

	// TODO:
	User map[string]any `json:"user"`

	// TODO:
	UserPlanLimits map[string]any `json:"user_plan_limits"`

	// TODO:
	UserSettings map[string]any `json:"user_settings"`
}

func (r *Response) UnmarshalJSON(data []byte) error {
	type alias Response
	aux := &struct {
		SyncStatus map[uuid.UUID]any `json:"sync_status"`
		*alias
	}{alias: (*alias)(r)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.SyncStatus = make(map[uuid.UUID]error, len(aux.SyncStatus))
	for k, v := range aux.SyncStatus {
		switch v := v.(type) {
		case string:
			if v == syncStatusOk {
				r.SyncStatus[k] = nil
			} else {
				r.SyncStatus[k] = errors.New(strings.ToLower(v))
			}
		case map[string]any:
			if err, ok := v["error"]; ok {
				if err, ok := err.(string); ok {
					r.SyncStatus[k] = errors.New(strings.ToLower(err))
				}
			} else {
				r.SyncStatus[k] = errors.New("unknown error")
			}
		default:
			r.SyncStatus[k] = errors.New("unknown error")
		}
	}

	return nil
}
