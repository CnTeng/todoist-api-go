package sync

// Delete a task and all its sub-tasks.
type ItemDeleteArgs struct {
	// ID of the task to delete.
	ID string `json:"id"` // required
}

func NewItemDeleteCommand(args ItemDeleteArgs) *Command {
	return NewCommand("item_delete", args)
}
