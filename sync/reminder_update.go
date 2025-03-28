package sync

// Update a reminder from the user account.
//
// See [Update a reminder] for more details.
//
// [Update a reminder]: https://developer.todoist.com/sync/v9/#update-a-reminder
type ReminderUpdateArgs struct {
	// Required.
	// The ID of the reminder.
	ID string `json:"id"`

	// Optional.
	// The user ID which should be notified of the reminder, typically the current
	// user ID creating the reminder.
	NotifyUID *string `json:"notify_uid,omitempty"`

	// Required.
	// The type of the reminder:
	//
	//   - relative: a time-based reminder specified in minutes from now.
	//   - absolute: a time-based reminder with a specific time and date in the future.
	//   - location: a location-based reminder.
	ReminderType ReminderType `json:"type"`

	// Optional.
	// The due date of the reminder. Reminders only support due dates with time,
	// since full-day reminders don't make sense.
	Due *Due `json:"due,omitempty"`

	// Optional.
	// The relative time in minutes before the due date of the item, in which the
	// reminder should be triggered. Note that the item should have a due date
	// with time set in order to add a relative reminder.
	MinuteOffset int `json:"minute_offset,omitempty"`

	// Optional.
	// An alias name for the location.
	Name *string `json:"name,omitempty"`

	// Optional.
	// The location latitude.
	LocLat *string `json:"loc_lat,omitempty"`

	// Optional.
	// The location longitude.
	LocLong *string `json:"loc_long,omitempty"`

	// Optional.
	// What should trigger the reminder:
	//
	//   - on_enter: entering the location.
	//   - on_leave: leaving the location.
	LocTrigger *string `json:"loc_trigger,omitempty"`

	// Optional.
	// The radius around the location that is still considered as part of the
	// location (in meters).
	Radius *int `json:"radius,omitempty"`
}

// Return "reminder_update" as command type.
func (args *ReminderUpdateArgs) Type() string {
	return "reminder_update"
}
