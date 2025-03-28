package sync

// Updates section attributes.
//
// See [Update a section] for more details.
//
// [Update a section]: https://developer.todoist.com/sync/v9/#update-a-section
type SectionUpdateArgs struct {
	// Required.
	// The ID of the section.
	ID string `json:"id"`

	// Optional.
	// The name of the section.
	Name *string `json:"name,omitempty"`

	// Optional.
	// Whether the section's tasks are collapsed.
	Collapsed *bool `json:"collapsed,omitempty"`
}

// Return "section_update" as command type.
func (args *SectionUpdateArgs) Type() string {
	return "section_update"
}
