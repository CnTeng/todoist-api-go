package sync

// Update child_order properties of items in bulk. See [Reorder items] for more
// details.
//
// [Reorder items]: https://developer.todoist.com/sync/v9#reorder-items
type ItemReorderArgs struct {
	// Required.
	// An array of objects to update.
	Items []struct {
		// The ID of the task.
		ID string `json:"id"`

		// The new order.
		ChildOrder int `json:"child_order"`
	}
}

// Return "item_reorder" as command type.
func (args *ItemReorderArgs) Type() string {
	return "item_reorder"
}
