package sync

type ProjectLeaveArgs struct {
	// Required.
	// The ID (v2_id) of the project to leave. It only accepts v2_id as the id.
	ProjectID string `json:"project_id"`
}

func (args *ProjectLeaveArgs) command() string {
	return "project_leave"
}
