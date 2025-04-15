package sync

// Completes a task and its sub-tasks and moves them to the archive. See also
// [TaskCloseArgs] for a simplified version of the command.
//
// See [Complete task] for more details.
//
// [Complete task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Complete-task
type TaskCompleteArgs struct {
	// Required.
	// Task ID to complete.
	ID string `json:"id"`

	// Optional.
	// RFC3339-formatted date of completion of the task (in UTC). If not set, the
	// server will set the value to the current timestamp.
	DateCompleted *string `json:"date_completed,omitempty"`
}

// Return "item_complete" as command type.
func (args *TaskCompleteArgs) Type() string {
	return "item_complete"
}

// Complete a recurring task. The reason why this is a special case is because
// we need to mark a recurring completion (and using item_update won't do this).
// See also [TaskCloseArgs] for a simplified version of the command.
//
// See [Complete a recurring task] for more details.
//
// [Complete a recurring task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Complete-a-recurring-task
type TaskCompleteRecurringArgs struct {
	// Required.
	// The ID of the item to update (a number or a temp id).
	ID string `json:"id"` // required

	// Optional.
	// The due date of the task. See the
	// https://todoist.com/api/v1/docs#tag/Due-dates section for more details.
	Due *Due `json:"due,omitempty"`

	// Optional.
	// Whether the task is completed or not. Defaults to true.
	IsForward *bool `json:"is_forward,omitempty"`

	// Optional.
	// Whether the subtasks are reset or not. Defaults to false.
	ResetSubtasks *bool `json:"reset_subtasks,omitempty"`
}

// Return "item_update_date_complete" as command type.
func (args *TaskCompleteRecurringArgs) Type() string {
	return "item_update_date_complete"
}
