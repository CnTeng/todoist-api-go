package sync

// Update personal label attributes.
//
// See [Update a personal label] for more details.
//
// [Update a personal label]: https://developer.todoist.com/sync/v9/#update-a-personal-label
type LabelUpdateArgs struct {
	// Required.
	// The ID of the label.
	ID string `json:"id"`

	// Optional.
	// The name of the label.
	Name *string `json:"name,omitempty"`

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

// Return "label_update" as command type.
func (args *LabelUpdateArgs) Type() string {
	return "label_update"
}
