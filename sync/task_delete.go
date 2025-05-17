package sync

type TaskDeleteArgs struct {
	// Required.
	// ID of the task to delete.
	ID string `json:"id"`
}

func (args *TaskDeleteArgs) command() string {
	return "item_delete"
}
