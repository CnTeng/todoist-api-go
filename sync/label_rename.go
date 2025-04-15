package sync

// Rename a shared label.
//
// See [Rename a shared label] for more details.
//
// [Rename a shared label]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Rename-a-shared-label
type LabelRenameArgs struct {
	// Required.
	// The current name of the label to modify.
	NameOld string `json:"name_old"`

	// Required.
	// The new name for the label.
	NameNew string `json:"name_new"`
}

// Return "label_rename" as command type.
func (args *LabelRenameArgs) Type() string {
	return "label_rename"
}
