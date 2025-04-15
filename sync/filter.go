package sync

// Availability of filters functionality and the maximum number of saved filters
// are dependent on the current user plan. These values are indicated by the
// filters and max_filters properties of the user [plan limits object].
//
// See [Filters] for more details.
//
// [plan limits object]: https://todoist.com/api/v1/docs#tag/Sync/User/User-plan-limits
// [Filters]: https://todoist.com/api/v1/docs#tag/Sync/Filters
type Filter struct {
	// The ID of the filter.
	ID string `json:"id"`

	// The name of the filter.
	Name string `json:"name"`

	// The query to search for.
	// https://www.todoist.com/help/articles/introduction-to-filters-V98wIH can be
	// found in the Todoist help page.
	Query string `json:"query"`

	// The color of the filter icon. Refer to the name column in the
	// https://todoist.com/api/v1/docs#tag/Colors guide for more info.
	Color Color `json:"color"`

	// Filter’s order in the filter list (where the smallest value should place
	// the filter at the top).
	ItemOrder int `json:"item_order"`

	// Whether the filter is marked as deleted (a true or false value).
	IsDeleted bool `json:"is_deleted"`

	// Whether the filter is a favorite (a true or false value).
	IsFavorite bool `json:"is_favorite"`

	// Filters from a cancelled subscription cannot be changed. This is a
	// read-only attribute (a true or false value).
	IsFrozen bool `json:"is_frozen"`
}
