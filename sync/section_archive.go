package sync

type SectionArchiveArgs struct {
	// Required.
	// Section ID to archive.
	ID string `json:"id"`
}

func (args *SectionArchiveArgs) command() string {
	return "section_archive"
}
