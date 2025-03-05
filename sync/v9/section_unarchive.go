package sync

// Unarchive a section.
//
// See [Unarchive a section] for more details.
//
// [Unarchive a section]: https://developer.todoist.com/sync/v9/#unarchive-a-section
type SectionUnarchiveArgs struct {
	// Required.
	// The ID of the section.
	ID string `json:"id"`
}

// Return "section_unarchive" as command type.
func (args *SectionUnarchiveArgs) Type() string {
	return "section_unarchive"
}
