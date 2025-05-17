package sync

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

func (args *LabelDeleteArgs) command() string {
	return "label_delete"
}

type LabelDeleteSharedArgs struct {
	// Required.
	// The name of the label to remove.
	Name string `json:"name"`
}

func (args *LabelDeleteSharedArgs) command() string {
	return "label_delete_occurrences"
}
