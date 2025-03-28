package sync

// Filter represents a filter.
//
// See [Filters] for more details.
//
// [Filters]: https://developer.todoist.com/sync/v9/#filters
type Filter struct {
	// The ID of the filter.
	ID string `json:"id"`

	// The name of the filter.
	Name string `json:"name"`

	// The query to search for.
	Query string `json:"query"`

	// The color of the filter icon. See
	// https://developer.todoist.com/guides/#colors for a list of available.
	Color Color `json:"color"`

	// Filter’s order in the filter list.
	ItemOrder int `json:"item_order"`

	// Whether the filter is marked as deleted/
	IsDeleted bool `json:"is_deleted"`

	// Whether the filter is a favorite.
	IsFavorite bool `json:"is_favorite"`
}
