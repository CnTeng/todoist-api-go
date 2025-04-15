package sync

// Unarchive a project. No ancestors will be unarchived along with the
// unarchived project. Instead, the project is unarchived alone, loses any
// parent relationship (becomes a root project), and is placed at the end of the
// list of other root projects.
//
// See [Unarchive a project] for more details.
//
// [Unarchive a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Unarchive-a-project
type ProjectUnarchiveArgs struct {
	// Required.
	// ID of the project to unarchive (could be a temp id).
	ID string `json:"id"`
}

func (args *ProjectUnarchiveArgs) command() string {
	return "project_unarchive"
}
