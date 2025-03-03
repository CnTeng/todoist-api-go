package sync

// Complete a recurring task. The reason why this is a special case is because
// we need to mark a recurring completion (and using item_update won't do this).
// See [Complete a recurring task] for more details. See also item_close for a
// simplified version of the command.
//
// [Complete a recurring task]: https://developer.todoist.com/sync/v9#complete-a-recurring-task
type ItemCompleteRecurringArgs struct {
	// Required.
	// Task ID to complete.
	ID string `json:"id"` // required

	// Optional.
	// The due date of the task.
	Due *Due `json:"due,omitempty"`

	// Optional.
	// Whether the task is completed or not. Defaults to true.
	IsForward *bool `json:"is_forward,omitempty"`

	// Optional.
	// Whether the subtasks are reset or not. Defaults to false.
	ResetSubtasks *bool `json:"reset_subtasks,omitempty"`
}

// Return "item_update_date_complete" as command type.
func (args *ItemCompleteRecurringArgs) Type() string {
	return "item_update_date_complete"
}
