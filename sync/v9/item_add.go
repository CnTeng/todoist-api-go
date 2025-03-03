package sync

// Add a new task to a project. See [Add an item] for more details.
//
// [Add an item]: https://developer.todoist.com/sync/v9#add-an-item
type ItemAddArgs struct {
	// Required.
	// The text of the task. This value may contain markdown-formatted text and
	// hyperlinks. See
	// https://todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9 for
	// more details.
	Content string `json:"content"`

	// Optional.
	// A description for the task. This value may contain markdown-formatted text
	// and hyperlinks. See
	// https://todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9 for
	// more details.
	Description *string `json:"description,omitempty"`

	// Optional.
	// The ID of the project to add the task to (a number or a temp id). By
	// default the task is added to the user’s Inbox project.
	ProjectID *string `json:"project_id,omitempty"`

	// Optional.
	// The due date of the task.
	Due *Due `json:"due,omitempty"`

	// Optional.
	// The deadline of the task.
	Deadline *Deadline `json:"deadline,omitempty"`

	// Optional.
	// The priority of the task, ranging from 1 (natural) to 4 (very urgent). Very
	// urgent is the priority 1 on clients. So, p1 will return 4 in the API.
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
	// The order of the task inside the Today or Next 7 days view. The smallest
	// value will place the task at the top.
	DayOrder *int `json:"day_order,omitempty"`

	// Optional.
	// Whether the task's sub-tasks are collapsed.
	Collapsed *bool `json:"collapsed,omitempty"`

	// Optional.
	// The task's labels. Represent either personal or shared labels.
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
	// item if it has a due date with time set. See
	// https://developer.todoist.com/sync/v9#user for more details about the
	// default reminder.
	AutoReminder *bool `json:"auto_reminder,omitempty"`

	// Optional.
	// When this option is enabled, the labels will be parsed from the task
	// content and added to the task. In case the label doesn't exist, a new one
	// will be created.
	AutoParseLabels *bool `json:"auto_parse_labels,omitempty"`

	// Optional.
	// The duration of the task.
	Duration *Duration `json:"duration,omitempty"`
}

// Return "item_add" as command type.
func (args *ItemAddArgs) Type() string {
	return "item_add"
}
