package sync

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

func (args *SectionAddArgs) command() string {
	return "section_add"
}
