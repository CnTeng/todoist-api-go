package sync

// Unarchive a project. No ancestors will be unarchived along with the
// unarchived project. Instead, the project is unarchived alone, loses any
// parent relationship (becomes a root project), and is placed at the end of the
// list of other root projects.
//
// See [Unarchive a project] for more details.
//
// [Unarchive a project]: https://developer.todoist.com/sync/v9/#unarchive-a-project
type ProjectUnarchiveArgs struct {
	// Required.
	// ID of the project to unarchive.
	ID string `json:"id"`
}

// Return "project_unarchive" as command type.
func (args *ProjectUnarchiveArgs) Type() string {
	return "project_unarchive"
}
