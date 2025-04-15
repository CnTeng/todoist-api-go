package sync

// Add a new task to a project.
//
// See [Add a task] for more details.
//
// [Add a task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Add-a-task
type TaskAddArgs struct {
	// Required.
	// The text of the task. This value may contain markdown-formatted text and
	// hyperlinks. Details on markdown support can be found in the
	// https://www.todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	// in the Help Center.
	Content string `json:"content"`

	// Optional.
	// A description for the task. This value may contain markdown-formatted text
	// and hyperlinks. Details on markdown support can be found in the
	// https://www.todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	// in the Help Center.
	Description *string `json:"description,omitempty"`

	// Optional.
	// The ID of the project to add the task to (a number or a temp id). By
	// default the task is added to the user’s Inbox project.
	ProjectID *string `json:"project_id,omitempty"`

	// Optional.
	// The due date of the task. See the
	// https://todoist.com/api/v1/docs#tag/Due-dates section for more details.
	Due *Due `json:"due,omitempty"`

	// Optional.
	// The deadline of the task. See the
	// https://todoist.com/api/v1/docs#tag/Deadlines section for more details.
	Deadline *Deadline `json:"deadline,omitempty"`

	// Optional.
	// The priority of the task (a number between 1 and 4, 4 for very urgent and
	// 1 for natural).
	Priority *int `json:"priority,omitempty"`

	// Optional.
	// The ID of the parent task. Set to null for root tasks.
	ParentID *string `json:"parent_id,omitempty"`

	// Optional.
	// The order of the task. Defines the position of the task among all the tasks
	// with the same parent.
	ChildOrder *int `json:"child_order,omitempty"`

	// Optional.
	// The ID of the parent section. Set to null for tasks not belonging to a
	// section.
	SectionID *string `json:"section_id,omitempty"`

	// Optional.
	// The order of the task inside the Today or Next 7 days view. (a number,
	// where the smallest value would place the task at the top).
	DayOrder *int `json:"day_order,omitempty"`

	// Optional.
	// Whether the task's sub-tasks are collapsed (a true or false value).
	Collapsed *bool `json:"collapsed,omitempty"`

	// Optional.
	// The task's labels (a list of names that may represent either personal or
	// shared labels).
	Labels []string `json:"labels,omitempty"`

	// Optional.
	// The ID of the user who assigned the task. This makes sense for shared
	// projects only. Accepts any user ID from the list of project collaborators.
	// If this value is unset or invalid, it will automatically be set up to your
	// uid.
	AssignedByUID *string `json:"assigned_by_uid,omitempty"`

	// Optional.
	// The ID of user who is responsible for accomplishing the current task. This
	// makes sense for shared projects only. Accepts any user ID from the list of
	// project collaborators or null or an empty string to unset.
	ResponsibleUID *string `json:"responsible_uid,omitempty"`

	// Optional.
	// When this option is enabled, the default reminder will be added to the new
	// item if it has a due date with time set. See also the
	// https://todoist.com/api/v1/docs#tag/Sync/User for more info about the
	// default reminder.
	AutoReminder *bool `json:"auto_reminder,omitempty"`

	// Optional.
	// When this option is enabled, the labels will be parsed from the task
	// content and added to the task. In case the label doesn't exist, a new one
	// will be created.
	AutoParseLabels *bool `json:"auto_parse_labels,omitempty"`

	// Optional.
	// The task's duration. Includes a positive integer (greater than zero) for
	// the amount of time the task will take, and the unit of time that the amount
	// represents which must be either minute or day. Both the amount and unit
	// must be defined.
	Duration *Duration `json:"duration,omitempty"`
}

// Return "item_add" as command type.
func (args *TaskAddArgs) Type() string {
	return "item_add"
}
