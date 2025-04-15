package sync

// Delete a personal label.
//
// See [Delete a personal label] for more details.
//
// [Delete a personal label]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Delete-a-personal-label
type LabelDeleteArgs struct {
	// Required.
	// The ID of the label.
	ID string `json:"id"`

	// Optional.
	// A string value, either all (default) or none. If no value or all is passed,
	// the personal label will be removed and any instances of the label will also
	// be removed from tasks (including tasks in shared projects). If none is
	// passed, the personal label will be removed from the user's account but it
	// will continue to appear on tasks as a shared label.
	Cascade *string `json:"cascade,omitempty"`
}

// Return "label_delete" as command type.
func (args *LabelDeleteArgs) Type() string {
	return "label_delete"
}

// Deletes all occurrences of a shared label from any active tasks.
//
// See [Delete shared label occurrences] for more details.
//
// [Delete shared label occurrences]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Delete-shared-label-occurrences
type LabelDeleteOccurrencesArgs struct {
	// Required.
	// The name of the label to remove.
	Name string `json:"name"`
}

// Return "label_delete_occurrences" as command type.
func (args *LabelDeleteOccurrencesArgs) Type() string {
	return "label_delete_occurrences"
}
