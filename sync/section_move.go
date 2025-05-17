package sync

type SectionMoveArgs struct {
	// Required.
	// The ID of the section.
	ID string `json:"id"`

	// Optional.
	// ID of the destination project.
	ProjectID *string `json:"project_id,omitempty"`
}

func (args *SectionMoveArgs) command() string {
	return "section_move"
}
