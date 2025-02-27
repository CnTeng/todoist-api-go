package sync

// Completes a task and its sub-tasks and moves them to the archive.
// See also item_close for a simplified version of the command.
type ItemCompleteArgs struct {
	// Task ID to complete.
	ID string `json:"id"` // required

	// RFC3339-formatted date of completion of the task (in UTC).
	// If not set, the server will set the value to the current timestamp.
	DateCompleted *string `json:"date_completed,omitempty"` // optional
}
