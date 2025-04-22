// Package sync defines the models of the sync API.
package sync

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type SyncRequest struct {
	// A special string, used to allow the client to perform incremental sync.
	// Pass * to retrieve all active resource data. More details about this below.
	SyncToken *string `json:"sync_token,omitempty"`

	// Used to specify what resources to fetch from the server. It should be a
	// JSON-encoded array of strings. Here is a list of available resource types:
	// labels, projects, items, notes, sections, filters, reminders,
	// reminders_location, locations, user, live_notifications, collaborators,
	// user_settings, notification_settings, user_plan_limits, completed_info,
	// stats, workspaces, workspace_users. You may use all to include all the
	// resource types. Resources can also be excluded by prefixing a - prior to
	// the name, for example, -projects
	ResourceTypes *ResourceTypes `json:"resource_types,omitempty"`

	// A JSON array of Command objects. Each command will be processed in the
	// specified order.
	Commands Commands `json:"commands,omitempty"`
}

const (
	DefaultSyncToken = "*"
	syncStatusOk     = "ok"
)

type SyncErr struct {
	// Error code.
	ErrorCode int `json:"error_code"`

	// Error message.
	ErrorMsg string `json:"error"`

	// Error tag.
	ErrorTag string `json:"error_tag"`

	// HTTP status code.
	HTTPCode int `json:"http_code"`

	// Extra error information.
	ErrorExtra map[string]any `json:"error_extra"`
}

func (e *SyncErr) Error() string {
	if e == nil {
		return ""
	}
	return strings.ToLower(e.ErrorMsg)
}

type SyncResponse struct {
	// A new synchronization token. Used by the client in the next sync request to
	// perform an incremental sync.
	SyncToken string `json:"sync_token"`

	// Whether the response contains all data (a full synchronization) or just the
	// incremental updates since the last sync.
	FullSync bool `json:"full_sync"`

	// A dictionary object containing result of each command execution. The key
	// will be the command's uuid field and the value will be the result status of
	// the command execution.
	SyncStatus map[uuid.UUID]error `json:"sync_status"`

	// A dictionary object that maps temporary resource IDs to real resource IDs.
	TempIDMapping map[uuid.UUID]string `json:"temp_id_mapping"`

	// TODO:
	// A user object.
	User map[string]any `json:"user,omitempty"`

	// An array of project objects.
	Projects []*Project `json:"projects,omitempty"`

	// An array of task objects.
	Tasks []*Task `json:"items,omitempty"`

	// An array of task comments objects.
	TaskComments []any `json:"notes,omitempty"`

	// An array of project comments objects.
	ProjectComments []any `json:"project_notes,omitempty"`

	// An array of section objects.
	Sections []*Section `json:"sections,omitempty"`

	// An array of personal label objects.
	Labels []*Label `json:"labels,omitempty"`

	// An array of filter objects.
	Filters []*Filter `json:"filters,omitempty"`

	// A JSON object specifying the order of items in daily agenda.
	DayOrders map[string]int `json:"day_orders,omitempty"`

	// An array of reminder objects.
	Reminders []*Reminder `json:"reminders,omitempty"`

	// An array of location objects.
	Locations []*Location `json:"locations,omitempty"`

	// A JSON object containing all collaborators for all shared projects. The
	// projects field contains the list of all shared projects, where the user
	// acts as one of collaborators.
	Collaborators []any `json:"collaborators,omitempty"`

	// An array specifying the state of each collaborator in each project. The
	// state can be invited, active, inactive, deleted.
	CollaboratorStates []any `json:"collaborator_states,omitempty"`

	// An array of live_notification objects.
	LiveNotifications []any `json:"live_notifications,omitempty"`

	// What is the last live notification the user has seen? This is used to
	// implement unread notifications.
	LiveNotificationsLastReadID string `json:"live_notifications_last_read_id"`

	// TODO:
	// A JSON object containing user settings.
	UserSettings map[string]any `json:"user_settings,omitempty"`

	// TODO:
	// A JSON object containing user plan limits.
	UserPlanLimits map[string]any `json:"user_plan_limits,omitempty"`

	// A JSON object containing workspace objects.
	Workspaces []any `json:"workspaces,omitempty"`

	// A JSON object containing workspace_user objects. Only in incremental sync.
	WorkspaceUsers []any `json:"workspace_users,omitempty"`
}

func (r *SyncResponse) UnmarshalJSON(data []byte) error {
	type alias SyncResponse
	aux := &struct {
		SyncStatus map[uuid.UUID]json.RawMessage `json:"sync_status"`
		*alias
	}{alias: (*alias)(r)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.SyncStatus = make(map[uuid.UUID]error, len(aux.SyncStatus))
	for k, v := range aux.SyncStatus {
		var status string
		if json.Unmarshal(v, &status) == nil && status == syncStatusOk {
			r.SyncStatus[k] = nil
			continue
		}

		syncErr := &SyncErr{}
		if err := json.Unmarshal(v, syncErr); err != nil {
			syncErr.ErrorMsg = fmt.Sprintf("unknown error: %s", string(v))
		}

		r.SyncStatus[k] = syncErr
	}

	return nil
}
