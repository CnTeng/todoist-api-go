package sync

// Delete a task and all its sub-tasks.
//
// See [Delete a task] for more details.
//
// [Delete a task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Delete-tasks
type TaskDeleteArgs struct {
	// Required.
	// ID of the task to delete.
	ID string `json:"id"`
}

// Return "item_delete" as command type.
func (args *TaskDeleteArgs) Type() string {
	return "item_delete"
}
