package sync

// Update parent project relationships of the project. See [Move a project] for
// more details.
//
// [Move a project]: https://developer.todoist.com/sync/v9/#move-a-project
type ProjectMoveArgs struct {
	// Required.
	// The ID of the project.
	ID string `json:"id"`

	// Required.
	// The ID of the parent project (could be temp id). If set to null, the
	// project will be moved to the root
	ParentID string `json:"parent_id"`
}

// Return "project_move" as command type.
func (args *ProjectMoveArgs) Type() string {
	return "project_move"
}
