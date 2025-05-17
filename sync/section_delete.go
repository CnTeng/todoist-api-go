package sync

type SectionDeleteArgs struct {
	// Required.
	// The ID of the section.
	ID string `json:"id"`
}

func (args *SectionDeleteArgs) command() string {
	return "section_delete"
}
