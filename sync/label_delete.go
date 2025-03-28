package sync

// Delete a personal label.
//
// See [Delete a personal label] for more details.
//
// [Delete a personal label]: https://developer.todoist.com/sync/v9/#delete-a-personal-label
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
