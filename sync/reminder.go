package sync

// Reminder represents a reminder in Todoist.
//
// See [Reminders] for more details.
//
// [Reminders]: https://todoist.com/api/v1/docs#tag/Sync/Reminders
type Reminder struct {
	// The ID of the reminder.
	ID string `json:"id"`

	// The user ID which should be notified of the reminder, typically the current
	// user ID creating the reminder.
	NotifyUID string `json:"notify_uid"`

	// The item ID for which the reminder is about.
	ItemID string `json:"item_id"`

	// The type of the reminder:
	//
	//   - relative: a time-based reminder specified in minutes from now.
	//   - absolute: a time-based reminder with a specific time and date in the future.
	//   - location: a location-based reminder.
	ReminderType ReminderType `json:"type"`

	// The due date of the reminder. See the
	// https://todoist.com/api/v1/docs#tag/Due-dates section for more details.
	// Note that reminders only support due dates with time, since full-day
	// reminders don't make sense.
	Due Due `json:"due"`

	// The relative time in minutes before the due date of the item, in which the
	// reminder should be triggered. Note that the item should have a due date
	// with time set in order to add a relative reminder.
	MinuteOffset int `json:"minute_offset"`

	// An alias name for the location.
	Name *string `json:"name"`

	// The location latitude.
	LocLat *string `json:"loc_lat"`

	// The location longitude.
	LocLong *string `json:"loc_long"`

	// What should trigger the reminder:
	//
	//   - on_enter: entering the location.
	//   - on_leave: leaving the location.
	LocTrigger *string `json:"loc_trigger"`

	// The radius around the location that is still considered as part of the
	// location (in meters).
	Radius *int `json:"radius"`

	// Whether the reminder is marked as deleted.
	IsDeleted bool `json:"is_deleted"`
}

// ReminderType represents the type of the reminder.
type ReminderType string

const (
	// A location-based reminder.
	RelativeReminder ReminderType = "relative"

	// A time-based reminder with a specific time and date in the future.
	AbsoluteReminder ReminderType = "absolute"

	// A time-based reminder specified in minutes from now.
	LocationReminder ReminderType = "location"
)
