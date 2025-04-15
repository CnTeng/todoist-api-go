package sync

// Uncomplete and restore an archived item.
//
// Any ancestor items or sections will also be reinstated. Items will have the
// checked value reset.
//
// The reinstated items and sections will appear at the end of the list within
// their parent, after any previously active tasks.
//
// See [Uncomplete task] for more details.
//
// [Uncomplete task]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Uncomplete-item
type TaskUncompleteArgs struct {
	// Required.
	// Task ID to uncomplete
	ID string `json:"id"`
}

// Return "item_uncomplete" as command type.
func (args *TaskUncompleteArgs) Type() string {
	return "item_uncomplete"
}
