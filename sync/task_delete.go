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

func (args *TaskDeleteArgs) command() string {
	return "item_delete"
}
