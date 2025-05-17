package sync

type ProjectUnarchiveArgs struct {
	// Required.
	// ID of the project to unarchive (could be a temp id).
	ID string `json:"id"`
}

func (args *ProjectUnarchiveArgs) command() string {
	return "project_unarchive"
}
