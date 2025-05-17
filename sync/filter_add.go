package sync

type FilterAddArgs struct {
	// Required.
	// The name of the filter.
	Name string `json:"name"`

	// Required.
	// The query to search for.
	// https://www.todoist.com/help/articles/introduction-to-filters-V98wIH can be
	// found in the Todoist help page.
	Query string `json:"query"`

	// Optional.
	// The color of the filter icon. Refer to the name column in the
	// https://todoist.com/api/v1/docs#tag/Colors guide for more info.
	Color *Color `json:"color,omitempty"`

	// Optional.
	// Filter’s order in the filter list (where the smallest value should place
	// the filter at the top).
	ItemOrder *int `json:"item_order,omitempty"`

	// Optional.
	// Whether the filter is a favorite (a true or false value).
	IsFavorite *bool `json:"is_favorite,omitempty"`
}

func (args *FilterAddArgs) command() string {
	return "filter_add"
}
