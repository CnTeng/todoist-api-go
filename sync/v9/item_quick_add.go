package sync

// Add a new item using the Quick Add implementation similar to that used in the
// official clients. See [Quick add an item] for more details.
//
// [Quick add an item]: https://developer.todoist.com/sync/v9#quick-add-an-item
type ItemQuickAddParams struct {
	// Required.
	// The text of the task that is parsed. It can include a due date in free form
	// text, a project name starting with the # character (without spaces), a
	// label starting with the @ character, an assignee starting with the +
	// character, a priority (e.g., p1), or a description starting from // until
	// the end of the text.
	Text string `json:"text" url:"text"`

	// Optional.
	// The content of the note.
	Note *string `json:"note,omitempty" url:"note,omitempty"`

	// Optional.
	// The date of the reminder, added in free form text.
	Reminder *string `json:"reminder,omitempty" url:"reminder,omitempty"`

	// Optional.
	// When this option is enabled, the default reminder will be added to the new
	// item if it has a due date with time set. See
	// https://developer.todoist.com/sync/v9#user for more details about the
	// default reminder.
	AutoReminder *bool `json:"auto_reminder,omitempty" url:"auto_reminder,omitempty"`
}
