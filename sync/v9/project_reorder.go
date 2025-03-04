package sync

// Update child_order properties of projects in bulk. See [Reorder projects] for
// more details.
//
// [Reorder projects]: https://developer.todoist.com/sync/v9/#reorder-projects
type ProjectReorderArgs struct {
	// Required.
	// An array of objects to update.
	Projects []struct {
		// The ID of the project.
		ID string `json:"id"`

		// The new order.
		ChildOrder int `json:"child_order"`
	}
}

// Return "project_reorder" as command type.
func (args *ProjectReorderArgs) Type() string {
	return "project_reorder"
}
