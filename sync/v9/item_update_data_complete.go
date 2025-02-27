package sync

type ItemUpdateDataCompleteArgs struct {
	// The ID of the item to update (a number or a temp id).
	ID string `json:"id"` // required

	// The due date of the task.
	//
	// See the [Due] for more details.
	Due *Due `json:"due,omitempty"` // optional

	// Set this argument to 1 for completion, or 0 for uncompletion (e.g., via undo).
	// By default, this argument is set to 1 (completion).
	IsForward *bool `json:"is_forward,omitempty"` // optional

	// Set this property to 1 to reset subtasks when a recurring task is completed.
	// By default, this property is not set (0), and subtasks will retain their
	// existing status when the parent task recurs.
	RestSubTasks *bool `json:"rest_subtasks,omitempty"` // optional
}

func NewItemUpdateDataCompleteCommand(args ItemUpdateDataCompleteArgs) *Command {
	return NewCommand("item_update_data_complete", args)
}
