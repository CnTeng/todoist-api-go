package sync

// Update task attributes. Please note that updating the parent, moving,
// completing or uncompleting tasks is not supported by item_update, more
// specific commands have to be used instead.
//
// See [Update a task] for more details.
//
// [Update a task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Update-a-task
type TaskUpdateArgs struct {
	// Required.
	// The ID of the task.
	ID string `json:"id"`

	// Optional.
	// A description for the task. This value may contain markdown-formatted text
	// and hyperlinks. Details on markdown support can be found in the
	// https://www.todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	// in the Help Center.
	Content *string `json:"content,omitempty"`

	// Optional.
	// A description for the task. This value may contain markdown-formatted text
	// and hyperlinks. Details on markdown support can be found in the
	// https://www.todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	// in the Help Center.
	Description *string `json:"description,omitempty"`

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
	// The order of the task inside the Today or Next 7 days view. (a number,
	// where the smallest value would place the task at the top).
	DayOrder *int `json:"day_order,omitempty"`

	// Optional.
	// The task's duration. Includes a positive integer (greater than zero) for
	// the amount of time the task will take, and the unit of time that the amount
	// represents which must be either minute or day. Both the amount and unit
	// must be defined.
	Duration *Duration `json:"duration,omitempty"`
}

// Return "task_update" as command type.
func (args *TaskUpdateArgs) Type() string {
	return "item_update"
}
