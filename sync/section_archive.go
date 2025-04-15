package sync

// Archive a section and all its child tasks.
//
// See [Archive a section] for more details.
//
// [Archive a section]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Archive-a-section
type SectionArchiveArgs struct {
	// Required.
	// Section ID to archive.
	ID string `json:"id"`
}

// Return "section_archive" as command type.
func (args *SectionArchiveArgs) Type() string {
	return "section_archive"
}
