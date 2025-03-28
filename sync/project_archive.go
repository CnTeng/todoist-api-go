package sync

// Archive a project and its descendants. See [Archive a project] for more
// details.
//
// [Archive a project]: https://developer.todoist.com/sync/v9/#archive-a-project
type ProjectArchiveArgs struct {
	// Required.
	// ID of the project to archive.
	ID string `json:"id"`
}

// Return "project_archive" as command type.
func (args *ProjectArchiveArgs) Type() string {
	return "project_archive"
}
