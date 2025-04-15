package rest

// Add a new task using the Quick Add implementation similar to that used in the
// official clients.
//
// See [Quick add a task] for more details.
//
// [Quick add a task]: https://todoist.com/api/v1/docs#tag/Tasks/operation/quick_add_api_v1_tasks_quick_post
type TaskQuickAddRequest struct {
	// Required.
	// The text of the task that is parsed. It can include a due date in free form
	// text, a project name starting with the # character (without spaces), a
	// label starting with the @ character, an assignee starting with the +
	// character, a priority (e.g., p1), a deadline between {} (e.g. {in 3 days}),
	// or a description starting from // until the end of the text.
	Text string `json:"text" url:"text"`

	// Optional.
	// The note of the task.
	Note *string `json:"note,omitempty" url:"note,omitempty"`

	// Optional.
	// The reminder date in free form text.
	Reminder *string `json:"reminder,omitempty" url:"reminder,omitempty"`

	// Optional.
	// When this option is enabled, the default reminder will be added to the new
	// item if it has a due date with time set. See also
	// https://todoist.com/api/v1/docs#tag/Sync/User for more info about the
	// default reminder.
	AutoReminder bool `json:"auto_reminder" url:"auto_reminder"`

	// Optional.
	// The meta of the task.
	Meta bool `json:"meta" url:"meta"`
}
