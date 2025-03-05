package sync

// Rename a shared label.
//
// See [Rename a shared label] for more details.
//
// [Rename a shared label]: https://developer.todoist.com/sync/v9/#rename-a-shared-label
type LabelRenameSharedArgs struct {
	// Required.
	// The current name of the label to modify.
	NameOld string `json:"name_old"`

	// Required.
	// The new name for the label.
	NameNew string `json:"name_new"`
}

// Return "label_rename" as command type.
func (args *LabelRenameSharedArgs) Type() string {
	return "label_rename"
}
