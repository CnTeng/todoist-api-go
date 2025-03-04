package sync

// Delete an existing project and all its descendants. See [Delete a project]
// for more details.
//
// [Delete a project]: https://developer.todoist.com/sync/v9/#delete-a-project
type ProjectDeleteArgs struct {
	// Required.
	// ID of the project to delete.
	ID string `json:"id"`
}

// Return "project_delete" as command type.
func (args *ProjectDeleteArgs) Type() string {
	return "project_delete"
}
