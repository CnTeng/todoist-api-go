package sync

import (
	"encoding/json"
	"net/url"

	"github.com/google/uuid"
)

type Command struct {
	Type   string    `json:"type"`
	Args   any       `json:"args"`
	UUID   uuid.UUID `json:"uuid"`
	TempID uuid.UUID `json:"temp_id"`
}

func NewCommand(t string, args any) *Command {
	return &Command{
		Type:   t,
		Args:   args,
		UUID:   uuid.New(),
		TempID: uuid.New(),
	}
}

type Commands []*Command

func (cs *Commands) EncodeValues(key string, v *url.Values) error {
	b, err := json.Marshal(cs)
	if err != nil {
		return err
	}

	v.Add(key, string(b))
	return nil
}

type Response struct {
	SyncToken    string         `json:"sync_token"`
	FullSync     bool           `json:"full_sync"`
	Projects     []*Project     `json:"projects"`
	Items        []*Item        `json:"items"`
	ItemNotes    []*ItemNote    `json:"notes"`
	ProjectNotes []*ProjectNote `json:"project_notes"`
	Sections     []*Section     `json:"sections"`
	Labels       []*Label       `json:"labels"`
	Filters      []*Filter      `json:"filters"`
	Reminders    []*Reminder    `json:"reminders"`
}
