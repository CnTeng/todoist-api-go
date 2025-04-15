package sync

// Update child_order properties of projects in bulk.
//
// See [Reorder projects] for more details.
//
// [Reorder projects]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Reorder-projects
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

func (args *ProjectReorderArgs) command() string {
	return "project_reorder"
}
