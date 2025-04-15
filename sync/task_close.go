package sync

// A simplified version of item_complete / item_update_date_complete. The
// command does exactly what official clients do when you close a task:
//
//   - regular tasks are completed and moved to the archive
//   - recurring tasks are scheduled to their next occurrence
//
// See [Close task] for more details.
//
// [Close task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Close-task
type TaskCloseArgs struct {
	// Required.
	// The ID of the item to close (a number or a temp id).
	ID string `json:"id"`
}

func (args *TaskCloseArgs) command() string {
	return "item_close"
}
