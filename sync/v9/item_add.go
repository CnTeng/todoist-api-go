package sync

// Add a new task to a project.
type ItemAddArgs struct {
	// The text of the task.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the [Text Formatting article] in the Help Center.
	//
	// [Text Formatting article]: https://todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	Content string `json:"content"` // required

	// A description for the task.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the [Text Formatting article] in the Help Center.
	//
	// [Text Formatting article]: https://todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	Description *string `json:"description,omitempty"` // optional

	// The ID of the project to add the task to (a number or a temp id).
	// By default the task is added to the user’s Inbox project.
	ProjectID *string `json:"project_id,omitempty"` // optional

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

	// The ID of the parent task. Set to null for root tasks.
	ParentID *string `json:"parent_id,omitempty"` // optional

	// The order of the task. Defines the position of the task among all the tasks with the same parent.
	ChildOrder *int `json:"child_order,omitempty"` // optional

	// The ID of the parent section. Set to null for tasks not belonging to a section.
	SectionID *string `json:"section_id,omitempty"` // optional

	// The order of the task inside the Today or Next 7 days view (a number,
	// where the smallest value would place the task at the top).
	DayOrder *int `json:"day_order,omitempty"` // optional

	// Whether the task's sub-tasks are collapsed.
	Collapsed *bool `json:"collapsed,omitempty"` // optional

	// The task's labels (a list of names that may represent either personal or shared labels).
	Labels []string `json:"labels,omitempty"` // optional

	// The ID of the user who created the task. This makes sense for shared projects only.
	// For tasks created before 31 Oct 2019 the value is set to null.
	// Cannot be set explicitly or changed via API.
	AddedByUID *string `json:"added_by_uid,omitempty"` // optional

	// The ID of the user who assigned the task. This makes sense for shared projects only.
	// Accepts any user ID from the list of project collaborators.
	// If this value is unset or invalid, it will automatically be set up to your uid.
	AssignedByUID *string `json:"assigned_by_uid,omitempty"` // optional

	// The ID of user who is responsible for accomplishing the current task.
	// This makes sense for shared projects only.
	// Accepts any user ID from the list of project collaborators or null or an empty string to unset.
	ResponsibleUID *string `json:"responsible_uid,omitempty"` // optional

	// Whether the task is marked as completed.
	Checked *bool `json:"checked,omitempty"` // optional

	// Whether the task is marked as deleted.
	IsDeleted *bool `json:"is_deleted,omitempty"` // optional

	// Identifier to find the match between tasks in shared projects of different collaborators.
	// When you share a task, its copy has a different ID in the projects of your collaborators.
	// To find a task in another account that matches yours, you can use the "sync_id" attribute.
	// For non-shared tasks, the attribute is null.
	SyncID *string `json:"sync_id,omitempty"` // optional

	// The date when the task was completed (or null if not completed).
	CompletedAt *string `json:"completed_at,omitempty"` // optional

	// The date when the task was created.
	AddedAt *string `json:"added_at,omitempty"` // optional

	// Object representing a task's duration.
	//
	// See the [Duration] for more details.
	Duration *Duration `json:"duration,omitempty"` // optional
}

func NewItemAddCommand(args *ItemAddArgs) *Command {
	return NewCommand("item_add", args)
}
