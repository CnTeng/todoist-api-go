package sync

// Delete an existing project and all its descendants. Workspace projects can
// only be deleted by ADMINs and it must be archived first.
//
// See [Delete a project] for more details.
//
// [Delete a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Delete-a-project
type ProjectDeleteArgs struct {
	// Required.
	// ID of the project to delete (could be a temp id).
	ID string `json:"id"`
}

func (args *ProjectDeleteArgs) command() string {
	return "project_delete"
}
