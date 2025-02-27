package sync

// Filter represents a filter in the Todoist application.
type Filter struct {
	// ID is the ID of the filter.
	ID string `json:"id"`

	// Name is the name of the filter.
	Name string `json:"name"`

	// Query is the query to search for. Examples of searches can be found in the Todoist help page.
	Query string `json:"query"`

	// Color is the color of the filter icon.
	Color string `json:"color"`

	// ItemOrder is the filter’s order in the filter list.
	ItemOrder int `json:"item_order"`

	// IsDeleted indicates whether the filter is marked as deleted.
	IsDeleted bool `json:"is_deleted"`

	// IsFavorite indicates whether the filter is a favorite.
	IsFavorite bool `json:"is_favorite"`
}
