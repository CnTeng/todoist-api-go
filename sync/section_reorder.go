package sync

// Update section_order properties of sections in bulk.
//
// See [Reorder sections] for more details.
//
// [Reorder sections]: https://developer.todoist.com/sync/v9/#reorder-sections
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

// Return "section_reorder" as command type.
func (args *SectionReorderArgs) Type() string {
	return "section_reorder"
}
