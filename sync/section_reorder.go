package sync

type SectionReorderArgs struct {
	// Required.
	// An array of objects to update.
	Sections []struct {
		// The ID of the project.
		ID string `json:"id"`

		// The new order.
		SectionOrder int `json:"section_order"`
	}
}

func (args *SectionReorderArgs) command() string {
	return "section_reorder"
}
