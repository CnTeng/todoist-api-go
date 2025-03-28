package sync

// Archive a section and all its child tasks.
//
// See [Archive a section] for more details.
//
// [Archive a section]: https://developer.todoist.com/sync/v9/#archive-a-section
type SectionArchiveArgs struct {
	// Required.
	// The ID of the section.
	ID string `json:"id"`
}

// Return "section_archive" as command type.
func (args *SectionArchiveArgs) Type() string {
	return "section_archive"
}
