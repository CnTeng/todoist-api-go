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

func NewCommand(args any) *Command {
	var t string
	switch args.(type) {
	case ItemAddArgs:
		t = "item_add"
	case ItemCloseArgs:
		t = "item_close"
	case ItemCompleteArgs:
		t = "item_complete"
	case ItemDeleteArgs:
		t = "item_delete"
	case ItemMoveArgs:
		t = "item_move"
	case ItemReorderArgs:
		t = "item_reorder"
	case ItemUncompleteArgs:
		t = "item_uncomplete"
	case ItemUpdateArgs:
		t = "item_update"
	case ItemUpdateDateCompleteArgs:
		t = "item_update_date_complete"
	}

	return &Command{
		Type:   t,
		Args:   args,
		UUID:   uuid.New(),
		TempID: uuid.New(),
	}
}

func NewCommandWithTempID(args any, tempID uuid.UUID) *Command {
	c := NewCommand(args)
	c.TempID = tempID
	return c
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
