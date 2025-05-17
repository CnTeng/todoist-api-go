package sync

type TaskUncompleteArgs struct {
	// Required.
	// Task ID to uncomplete
	ID string `json:"id"`
}

func (args *TaskUncompleteArgs) command() string {
	return "item_uncomplete"
}
