package sync

// Label represents a personal label in the Todoist application.
type Label struct {
	// ID is the ID of the label.
	ID string `json:"id"`

	// Name is the name of the label.
	Name string `json:"name"`

	// Color is the color of the label icon.
	Color string `json:"color"`

	// ItemOrder is the label’s order in the label list.
	ItemOrder int `json:"item_order"`

	// IsDeleted indicates whether the label is marked as deleted.
	IsDeleted bool `json:"is_deleted"`

	// IsFavorite indicates whether the label is a favorite.
	IsFavorite bool `json:"is_favorite"`
}
