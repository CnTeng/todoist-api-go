package sync

// Complete a task and its sub-tasks and move them to the archive. See
// [Complete item] for more details. See also item_close for a simplified
// version of the command.
//
// [Complete item]: https://developer.todoist.com/sync/v9#complete-item
type ItemCompleteArgs struct {
	// Required.
	// Task ID to complete.
	ID string `json:"id"`

	// Optional.
	// RFC3339-formatted date of completion of the task (in UTC). If not set, the
	// server will set the value to the current timestamp.
	DateCompleted *string `json:"date_completed,omitempty"`
}

// Return "item_complete" as command type.
func (args *ItemCompleteArgs) Type() string {
	return "item_complete"
}
