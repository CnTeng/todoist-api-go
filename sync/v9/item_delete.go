package sync

// Delete a task and all its sub-tasks. See [Delete item] for more details.
//
// [Delete item]: https://developer.todoist.com/sync/v9#delete-item
type ItemDeleteArgs struct {
	// Required.
	// ID of the task to delete.
	ID string `json:"id"`
}

// Return "item_delete" as command type.
func (args *ItemDeleteArgs) Type() string {
	return "item_delete"
}
