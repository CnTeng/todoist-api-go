package sync

// Update a filter.
//
// See [Update a filter] for more details.
//
// [Update a filter]: https://developer.todoist.com/sync/v9/#update-a-filter
type FilterUpdateArgs struct {
	// Required.
	// The ID of the filter.
	ID string `json:"id"`

	// Optional.
	// The name of the filter.
	Name *string `json:"name,omitempty"`

	// Optional.
	// The query to search for.
	Query *string `json:"query,omitempty"`

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

// Return "filter_update" as command type.
func (args *FilterUpdateArgs) Type() string {
	return "filter_update"
}
