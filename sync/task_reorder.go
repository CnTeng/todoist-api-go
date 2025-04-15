package sync

// Update child_order properties of items in bulk.
//
// See [Reorder tasks] for more details.
//
// [Reorder tasks]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Reorder-tasks
type TaskReorderArgs struct {
	// Required.
	// An array of objects to update.
	Items []struct {
		// The ID of the task.
		ID string `json:"id"`

		// The new order.
		ChildOrder int `json:"child_order"`
	}
}

func (args *TaskReorderArgs) command() string {
	return "item_reorder"
}
