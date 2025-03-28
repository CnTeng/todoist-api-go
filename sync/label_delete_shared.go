package sync

// Deletes all occurrences of a shared label from any active tasks.
//
// See [Delete shared label occurrences] for more details.
//
// [Delete shared label occurrences]: https://developer.todoist.com/sync/v9/#delete-shared-label-occurrences
type LabelDeleteSharedArgs struct {
	// Required.
	// The name of the label to remove.
	Name string `json:"name"`
}

// Return "label_delete_occurrences" as command type.
func (args *LabelDeleteSharedArgs) Type() string {
	return "label_delete_occurrences"
}
