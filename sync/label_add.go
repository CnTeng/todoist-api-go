package sync

// Add a new personal label.
//
// See [Add a personal label] for more details.
//
// [Add a personal label]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Add-a-personal-label
type LabelAddArgs struct {
	// Required.
	// The name of the label.
	Name string `json:"name"`

	// Optional.
	// The color of the label icon. Refer to the name column in the
	// https://todoist.com/api/v1/docs#tag/Colors guide for more info.
	Color *Color `json:"color,omitempty"`

	// Optional.
	// Label’s order in the label list (a number, where the smallest value should
	// place the label at the top).
	ItemOrder *int `json:"item_order,omitempty"`

	// Optional.
	// Whether the label is a favorite (a true or false value).
	IsFavorite *bool `json:"is_favorite,omitempty"`
}

func (args *LabelAddArgs) command() string {
	return "label_add"
}
