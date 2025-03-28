package sync

// Delete a reminder from the current user account.
//
// See [Delete a reminder] for more details.
//
// [Delete a reminder]: https://developer.todoist.com/sync/v9/#delete-a-reminder
type ReminderDeleteArgs struct {
	// Required.
	// The ID of the filter.
	ID string `json:"id"`
}

// Return "reminder_delete" as command type.
func (args *ReminderDeleteArgs) Type() string {
	return "reminder_delete"
}
