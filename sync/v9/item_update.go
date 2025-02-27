package sync

// Updates task attributes.
//
// Please note that updating the parent, moving, completing or uncompleting tasks is not supported by item_update,
// more specific commands have to be used instead.
type ItemUpdateArgs struct {
	// The ID of the task.
	ID string `json:"id"` // required

	// The text of the task.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the [Text Formatting article] in the Help Center.
	//
	// [Text Formatting article]: https://todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	Content *string `json:"content,omitempty"` // optional

	// A description for the task.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the [Text Formatting article] in the Help Center.
	//
	// [Text Formatting article]: https://todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	Description *string `json:"description,omitempty"` // optional

	// The due date of the task.
	//
	// See the [Due] for more details.
	Due *Due `json:"due,omitempty"` // optional

	// The deadline of the task.
	//
	// See the [Deadline] for more details.
	Deadline *Deadline `json:"deadline,omitempty"` // optional

	// The priority of the task (a number between 1 and 4, 4 for very urgent and 1 for natural).
	//
	// Note: Keep in mind that very urgent is the priority 1 on clients.
	// So, p1 will return 4 in the API.
	Priority *int `json:"priority,omitempty"` // optional

	// Whether the task's sub-tasks are collapsed.
	Collapsed *bool `json:"collapsed,omitempty"` // optional

	// The task's labels (a list of names that may represent either personal or shared labels).
	Labels []string `json:"labels,omitempty"` // optional

	// The ID of the user who assigned the task. This makes sense for shared projects only.
	// Accepts any user ID from the list of project collaborators.
	// If this value is unset or invalid, it will automatically be set up to your uid.
	AssignedByUID *string `json:"assigned_by_uid,omitempty"` // optional

	// The ID of user who is responsible for accomplishing the current task.
	// This makes sense for shared projects only.
	// Accepts any user ID from the list of project collaborators or null or an empty string to unset.
	ResponsibleUID *string `json:"responsible_uid,omitempty"` // optional

	// The order of the task inside the Today or Next 7 days view (a number,
	// where the smallest value would place the task at the top).
	DayOrder *int `json:"day_order,omitempty"` // optional

	// Object representing a task's duration.
	//
	// See the [Duration] for more details.
	Duration *Duration `json:"duration,omitempty"` // optional
}

func NewItemUpdateCommand(args ItemUpdateArgs) *Command {
	return NewCommand("item_update", args)
}
