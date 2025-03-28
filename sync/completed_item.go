package sync

import (
	"encoding/json"
	"time"
)

// CompletedTask represents a completed task entry. See [Get all completed items]
// for more details.
//
// [Get all completed items]: https://developer.todoist.com/sync/v9#get-all-completed-items
type CompletedTask struct {
	// The ID of the completed task entry.
	ID string `json:"id"`

	// The ID of the task.
	TaskID string `json:"task_id"`

	// The owner of the task.
	UserID string `json:"user_id"`

	// The ID of the parent project.
	ProjectID string `json:"project_id"`

	// The ID of the parent section.
	SectionID *string `json:"section_id"`

	// The text of the task. This value may contain markdown-formatted text and
	// hyperlinks. See
	// https://todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9 for
	// more details.
	Content string `json:"content"`

	// The date when the task was completed.
	CompletedAt time.Time `json:"completed_at"`

	// The number of notes of the task.
	NoteCount int `json:"note_count"`

	// Optional extra details.
	MetaData *string `json:"meta_data"`

	// If annotate_items is set to true, will contain the full item object.
	ItemObject *Item `json:"item_object"`
}

func (ct *CompletedTask) UnmarshalJSON(data []byte) error {
	type alias CompletedTask
	aux := &struct {
		CompletedAt string `json:"completed_at"`
		*alias
	}{alias: (*alias)(ct)}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	t, err := time.Parse(time.RFC3339, aux.CompletedAt)
	if err != nil {
		return err
	}
	ct.CompletedAt = t

	return nil
}
