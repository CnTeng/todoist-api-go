package sync

// Only used for teams.
//
// This command is used to remove self from a workspace project. As this is a
// v2-only command, it will only accept v2 project id.
//
// See [Leave a project] for more details.
//
// [Leave a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Leave-a-project
type ProjectLeaveArgs struct {
	// Required.
	// The ID (v2_id) of the project to leave. It only accepts v2_id as the id.
	ProjectID string `json:"project_id"`
}

func (args *ProjectLeaveArgs) command() string {
	return "project_leave"
}
