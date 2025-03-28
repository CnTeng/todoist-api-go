package sync

// Add a new personal label.
//
// See [Add a personal label] for more details.
//
// [Add a personal label]: https://developer.todoist.com/sync/v9/#add-a-personal-label
type LabelAddArgs struct {
	// Required.
	// The name of the label.
	Name string `json:"name"`

	// Optional.
	// The color of the label icon. See
	// https://developer.todoist.com/guides/#colors for a list of available.
	Color *Color `json:"color,omitempty"`

	// Optional.
	// The order of the Label in the label list.
	ItemOrder *int `json:"item_order,omitempty"`

	// Optional.
	// Whether the label is a favorite.
	IsFavorite *bool `json:"is_favorite,omitempty"`
}

// Return "label_add" as command type.
func (args *LabelAddArgs) Type() string {
	return "label_add"
}
