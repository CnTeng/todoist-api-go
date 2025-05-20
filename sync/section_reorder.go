package sync

type SectionReorderItem struct {
	// The ID of the project.
	ID string `json:"id"`

	// The new order.
	SectionOrder int `json:"section_order"`
}

type SectionReorderArgs struct {
	// Required.
	// An array of objects to update.
	Items []SectionReorderItem `json:"sections"`
}

func (args *SectionReorderArgs) command() string {
	return "section_reorder"
}
