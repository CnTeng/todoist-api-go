package sync

// Unarchive a section.
//
// See [Unarchive a section] for more details.
//
// [Unarchive a section]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Unarchive-a-section
type SectionUnarchiveArgs struct {
	// Required.
	// Section ID to unarchive
	ID string `json:"id"`
}

func (args *SectionUnarchiveArgs) command() string {
	return "section_unarchive"
}
