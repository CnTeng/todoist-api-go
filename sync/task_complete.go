package sync

type TaskCompleteArgs struct {
	// Required.
	// Task ID to complete.
	ID string `json:"id"`

	// Optional.
	// RFC3339-formatted date of completion of the task (in UTC). If not set, the
	// server will set the value to the current timestamp.
	DateCompleted *string `json:"date_completed,omitempty"`
}

func (args *TaskCompleteArgs) command() string {
	return "item_complete"
}

type TaskCompleteRecurringArgs struct {
	// Required.
	// The ID of the item to update (a number or a temp id).
	ID string `json:"id"` // required

	// Optional.
	// The due date of the task. See the
	// https://todoist.com/api/v1/docs#tag/Due-dates section for more details.
	Due *Due `json:"due,omitempty"`

	// Optional.
	// Whether the task is completed or not. Defaults to true.
	IsForward *bool `json:"is_forward,omitempty"`

	// Optional.
	// Whether the subtasks are reset or not. Defaults to false.
	ResetSubtasks *bool `json:"reset_subtasks,omitempty"`
}

func (args *TaskCompleteRecurringArgs) command() string {
	return "item_update_date_complete"
}
