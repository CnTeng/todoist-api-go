package sync

type ReminderDeleteArgs struct {
	// Required.
	// The ID of the reminder.
	ID string `json:"id"`
}

func (args *ReminderDeleteArgs) command() string {
	return "reminder_delete"
}
