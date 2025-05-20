package sync

type ProjectReorderItem struct {
	// The ID of the project.
	ID string `json:"id"`

	// The new order.
	ChildOrder int `json:"child_order"`
}

type ProjectReorderArgs struct {
	// Required.
	// An array of objects to update.
	Items []ProjectReorderItem `json:"projects"`
}

func (args *ProjectReorderArgs) command() string {
	return "project_reorder"
}
