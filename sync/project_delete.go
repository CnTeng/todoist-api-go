package sync

type ProjectDeleteArgs struct {
	// Required.
	// ID of the project to delete (could be a temp id).
	ID string `json:"id"`
}

func (args *ProjectDeleteArgs) command() string {
	return "project_delete"
}
