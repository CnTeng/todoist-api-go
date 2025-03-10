package sync

// Add a new section to a project.
//
// See [Add a section] for more details.
//
// [Add a section]: https://developer.todoist.com/sync/v9/#add-a-section
type SectionAddArgs struct {
	// Required.
	// The name of the section.
	Name string `json:"name"`

	// Required.
	// The ID of the parent project.
	ProjectID string `json:"project_id"`

	// Optional.
	// The order of the section. Defines the position of the section among all the
	// sections in the project.
	SectionOrder *int `json:"section_order,omitempty"`
}

// Return "section_add" as command type.
func (args *SectionAddArgs) Type() string {
	return "section_add"
}
