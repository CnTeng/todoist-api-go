package sync

// Move task to a different location.
//
// Only one of parent_id, section_id or project_id must be set.
//
// Note, to move an item from a section to no section, just use the project_id parameter,
// with the project it currently belongs to as a value.
type ItemMoveArgs struct {
	// The ID of the task.
	ID string `json:"id"` // required

	// The ID of the parent task. Set to null for root tasks.
	ParentID *string `json:"parent_id,omitempty"` // optional

	// The ID of the parent section. Set to null for tasks not belonging to a section.
	SectionID *string `json:"section_id,omitempty"` // optional

	// The ID of the project to add the task to (a number or a temp id).
	// By default the task is added to the user’s Inbox project.
	ProjectID *string `json:"project_id,omitempty"` // optional
}

func NewItemMoveCommand(args ItemMoveArgs) *Command {
	return NewCommand("item_move", args)
}
