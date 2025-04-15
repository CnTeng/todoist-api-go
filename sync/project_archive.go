package sync

// Archive a project and its descendants.
//
// See [Archive a project] for more details.
//
// [Archive a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Archive-a-project
type ProjectArchiveArgs struct {
	// Required.
	// ID of the project to archive (could be a temp id).
	ID string `json:"id"`
}

// Return "project_archive" as command type.
func (args *ProjectArchiveArgs) Type() string {
	return "project_archive"
}
