package sync

// Delete a section and all its child tasks.
//
// See [Delete a section] for more details.
//
// [Delete a section]: https://developer.todoist.com/sync/v9/#delete-a-section
type SectionDeleteArgs struct {
	// Required.
	// The ID of the section.
	ID string `json:"id"`
}

// Return "section_delete" as command type.
func (args *SectionDeleteArgs) Type() string {
	return "section_delete"
}
