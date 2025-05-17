package sync

type SectionUpdateArgs struct {
	// Required.
	// The ID of the section.
	ID string `json:"id"`

	// Optional.
	// The name of the section.
	Name *string `json:"name,omitempty"`

	// Optional.
	// Whether the section's tasks are collapsed (a true or false value).
	IsCollapsed *bool `json:"is_collapsed,omitempty"`
}

func (args *SectionUpdateArgs) command() string {
	return "section_update"
}
