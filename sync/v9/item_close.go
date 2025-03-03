package sync

// A simplified version of item_complete / item_update_date_complete. The
// command does exactly what official clients do when you close a task:
//
//   - regular tasks are completed and moved to the archive
//   - recurring tasks are scheduled to their next occurrence
//
// See [Close item] for more details.
//
// [Close item]: https://developer.todoist.com/sync/v9#close-item
type ItemCloseArgs struct {
	// Required.
	// The ID of the task.
	ID string `json:"id"`
}

// Return "item_close" as command type.
func (args *ItemCloseArgs) Type() string {
	return "item_close"
}
