package sync

// Update task attributes. See [Update an item] for more details.
//
// [Update an item]: https://developer.todoist.com/sync/v9#update-an-item
type ItemUpdateArgs struct {
	// Required.
	// The ID of the task.
	ID string `json:"id"`

	// Optional.
	// The text of the task. This value may contain markdown-formatted text and
	// hyperlinks. See
	// https://todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9 for
	// more details.
	Content *string `json:"content,omitempty"`

	// Optional.
	// A description for the task. This value may contain markdown-formatted text
	// and hyperlinks. See
	// https://todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9 for
	// more details.
	Description *string `json:"description,omitempty"`

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
	// The order of the task inside the Today or Next 7 days view. The smallest
	// value will place the task at the top.
	DayOrder *int `json:"day_order,omitempty"`

	// Optional.
	// The duration of the task.
	Duration *Duration `json:"duration,omitempty"`
}

// Return "task_update" as command type.
func (args *ItemUpdateArgs) Type() string {
	return "item_update"
}
