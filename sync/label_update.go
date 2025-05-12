package sync

// Update a personal label.
//
// See [Update a personal label] for more details.
//
// [Update a personal label]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Update-a-personal-label
type LabelUpdateArgs struct {
	// Required.
	// The ID of the label.
	ID string `json:"id"`

	// Optional.
	// The name of the label.
	Name *string `json:"name,omitempty"`

	// Optional.
	// The color of the label icon. Refer to the name column in the
	// https://todoist.com/api/v1/docs#tag/Colors guide for more info.
	Color *Color `json:"color,omitempty"`

	// Optional.
	// Label’s order in the label list.
	ItemOrder *int `json:"item_order,omitempty"`

	// Optional.
	// Whether the label is a favorite (a true or false value).
	IsFavorite *bool `json:"is_favorite,omitempty"`
}

func (args *LabelUpdateArgs) command() string {
	return "label_update"
}

// Rename a shared label.
//
// See [Rename a shared label] for more details.
//
// [Rename a shared label]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Rename-a-shared-label
type LabelRenameSharedArgs struct {
	// Required.
	// The current name of the label to modify.
	NameOld string `json:"name_old"`

	// Required.
	// The new name for the label.
	NameNew string `json:"name_new"`
}

func (args *LabelRenameSharedArgs) command() string {
	return "label_rename"
}
