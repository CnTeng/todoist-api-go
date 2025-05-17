package sync

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
