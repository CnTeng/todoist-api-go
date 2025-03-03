package sync

// Move task to a different location. See [Move an item] for more details.
//
// Only one of parent_id, section_id or parent_id must be set. To move an item
// out of a section, use the parent_id of its current project.
//
// [Move an item]: https://developer.todoist.com/sync/v9#move-an-item
type ItemMoveArgs struct {
	// Required.
	// The ID of the task.
	ID string `json:"id"`

	// Optional.
	// The ID of the parent task. The task becomes the last child task of the
	// parent task.
	ParentID *string `json:"parent_id,omitempty"`

	// Optional.
	// The ID of the parent section. The task becomes the last root task of the
	// section.
	SectionID *string `json:"section_id,omitempty"`

	// Optional.
	// The ID of the parent project. The task becomes the last root task of the
	// project.
	ProjectID *string `json:"project_id,omitempty"`
}

// Return "item_move" as command type.
func (args *ItemMoveArgs) Type() string {
	return "item_move"
}
