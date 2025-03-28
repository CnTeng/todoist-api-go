package sync

// Label represents a personal label.
//
// See [Labels] for more details.
//
// [Labels]: https://developer.todoist.com/sync/v9/#labels
type Label struct {
	// The ID of the label.
	ID string `json:"id"`

	// The name of the label.
	Name string `json:"name"`

	// The color of the label icon. See
	// https://developer.todoist.com/guides/#colors for a list of available.
	Color Color `json:"color"`

	// The order of the Label in the label list.
	ItemOrder int `json:"item_order"`

	// Whether the label is marked as deleted.
	IsDeleted bool `json:"is_deleted"`

	// Whether the label is a favorite.
	IsFavorite bool `json:"is_favorite"`
}
