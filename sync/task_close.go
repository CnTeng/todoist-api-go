package sync

type TaskCloseArgs struct {
	// Required.
	// The ID of the item to close (a number or a temp id).
	ID string `json:"id"`
}

func (args *TaskCloseArgs) command() string {
	return "item_close"
}
