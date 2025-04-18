package sync

// Delete a reminder from the current user account.
//
// See [Delete a reminder] for more details.
//
// [Delete a reminder]: https://todoist.com/api/v1/docs#tag/Sync/Reminders/Delete-a-reminder
type ReminderDeleteArgs struct {
	// Required.
	// The ID of the reminder.
	ID string `json:"id"`
}

func (args *ReminderDeleteArgs) command() string {
	return "reminder_delete"
}
