package sync

type ProjectArchiveArgs struct {
	// Required.
	// ID of the project to archive (could be a temp id).
	ID string `json:"id"`
}

func (args *ProjectArchiveArgs) command() string {
	return "project_archive"
}
