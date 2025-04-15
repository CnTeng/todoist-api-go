package sync

// Move section to a different location.
//
// See [Move a section] for more details.
//
// [Move a section]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Move-a-section
type SectionMoveArgs struct {
	// Required.
	// The ID of the section.
	ID string `json:"id"`

	// Optional.
	// ID of the destination project.
	ProjectID *string `json:"project_id,omitempty"`
}

// Return "section_move" as command type.
func (args *SectionMoveArgs) Type() string {
	return "section_move"
}
