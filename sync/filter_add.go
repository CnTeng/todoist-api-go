package sync

// Add a filter.
//
// See [Add a filter] for more details.
//
// [Add a filter]: https://developer.todoist.com/sync/v9/#add-a-filter
type FilterAddArgs struct {
	// Required.
	// The name of the filter.
	Name string `json:"name"`

	// Required.
	// The query to search for.
	Query string `json:"query"`

	// Optional.
	// The color of the filter icon. See
	// https://developer.todoist.com/guides/#colors for a list of available.
	Color *Color `json:"color,omitempty"`

	// Optional.
	// Filter’s order in the filter list.
	ItemOrder *int `json:"item_order,omitempty"`

	// Optional.
	// Whether the filter is a favorite.
	IsFavorite *bool `json:"is_favorite,omitempty"`
}

// Return "filter_add" as command type.
func (args *FilterAddArgs) Type() string {
	return "filter_add"
}
