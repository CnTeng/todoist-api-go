package sync

import "time"

// Task represents a task.
//
// See [Tasks] for more details.
//
// [Tasks]: https://todoist.com/api/v1/docs#tag/Sync/Tasks
type Task struct {
	// The ID of the task.
	ID string `json:"id"`

	// The owner of the task.
	UserID string `json:"user_id"`

	// The ID of the parent project.
	ProjectID string `json:"project_id"`

	// The text of the task. This value may contain markdown-formatted text and
	// hyperlinks. Details on markdown support can be found in the
	// https://www.todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	// in the Help Center.
	Content string `json:"content"`

	// A description for the task. This value may contain markdown-formatted text
	// and hyperlinks. Details on markdown support can be found in the
	// https://www.todoist.com/help/articles/format-text-in-a-todoist-task-e5dHw9
	// in the Help Center.
	Description string `json:"description"`

	// The due date of the task. See the
	// https://todoist.com/api/v1/docs#tag/Due-dates section for more details.
	Due *Due `json:"due"`

	// The deadline of the task. See the
	// https://todoist.com/api/v1/docs#tag/Deadlines section for more details.
	Deadline *Deadline `json:"deadline"`

	// The priority of the task (a number between 1 and 4, 4 for very urgent and
	// 1 for natural).
	Priority int `json:"priority"`

	// The ID of the parent task. Set to null for root tasks.
	ParentID *string `json:"parent_id"`

	// The order of the task. Defines the position of the task among all the tasks
	// with the same parent.
	ChildOrder int `json:"child_order"`

	// The ID of the parent section. Set to null for tasks not belonging to a
	// section.
	SectionID *string `json:"section_id"`

	// The order of the task inside the Today or Next 7 days view. (a number,
	// where the smallest value would place the task at the top).
	DayOrder int `json:"day_order"`

	// Whether the task's sub-tasks are collapsed (a true or false value).
	Collapsed bool `json:"collapsed"`

	// The task's labels (a list of names that may represent either personal or
	// shared labels).
	Labels []string `json:"labels"`

	// The ID of the user who created the task. This makes sense for shared
	// projects only. For tasks created before 31 Oct 2019 the value is set to
	// null. Cannot be set explicitly or changed via API.
	AddedByUID *string `json:"added_by_uid"`

	// The ID of the user who assigned the task. This makes sense for shared
	// projects only. Accepts any user ID from the list of project collaborators.
	// If this value is unset or invalid, it will automatically be set up to your
	// uid.
	AssignedByUID *string `json:"assigned_by_uid"`

	// The ID of user who is responsible for accomplishing the current task. This
	// makes sense for shared projects only. Accepts any user ID from the list of
	// project collaborators or null or an empty string to unset.
	ResponsibleUID *string `json:"responsible_uid"`

	// Whether the task is marked as completed (a true or false value).
	Checked bool `json:"checked"`

	// Whether the task is marked as deleted (a true or false value).
	IsDeleted bool `json:"is_deleted"`

	// The datetime when the task was created.
	AddedAt time.Time `json:"added_at"`

	// The datetime when the task was updated.
	UpdatedAt time.Time `json:"updated_at"`

	// The date when the task was completed.
	CompletedAt *string `json:"completed_at"`

	// The task's duration. Includes a positive integer (greater than zero) for
	// the amount of time the task will take, and the unit of time that the amount
	// represents which must be either minute or day. Both the amount and unit
	// must be defined.
	Duration *Duration `json:"duration"`
}
