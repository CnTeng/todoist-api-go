package sync

type TaskReorderItem struct {
	// The ID of the task.
	ID string `json:"id"`

	// The new order.
	ChildOrder int `json:"child_order"`
}

type TaskReorderArgs struct {
	// Required.
	// An array of objects to update.
	Items []TaskReorderItem `json:"items"`
}

func (args *TaskReorderArgs) command() string {
	return "item_reorder"
}
