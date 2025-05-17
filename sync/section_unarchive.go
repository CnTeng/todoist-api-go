package sync

type SectionUnarchiveArgs struct {
	// Required.
	// Section ID to unarchive
	ID string `json:"id"`
}

func (args *SectionUnarchiveArgs) command() string {
	return "section_unarchive"
}
