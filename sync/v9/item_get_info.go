package sync

// Get detailed information about an item and any related entities.
//
// A full sync only returns up to 10 notes for each item. If a client requires
// more, they can be downloaded with this endpoint.
//
// It returns a JSON object with the item, and details of any ancestors, the
// parent project, parent section if assigned, and any notes. See
// [ItemGetInfoResponse] for more details.
//
// See [Get item info] for more details.
//
// [Get item info]: https://developer.todoist.com/sync/v9#get-item-info
type ItemGetInfoParams struct {
	// Required.
	// The item's unique ID.
	ItemID string `json:"item_id" url:"item_id"`

	// Optional.
	// Whether to return ancestors, project, section and notes of the item. If it
	// is set to false, only the data for the item will be returned. Defaults to
	// true.
	AllData *bool `json:"all_data,omitempty" url:"all_data,omitempty"`
}

// [ItemGetInfoParams] returns a JSON object with the item, and details of any ancestors, the
// parent project, parent section if assigned, and any notes.
type ItemGetInfoResponse struct {
	Ancestors []*Item     `json:"ancestors"`
	Item      *Item       `json:"item"`
	Notes     []*ItemNote `json:"notes"`
	Project   *Project    `json:"project"`
	Section   *Section    `json:"section"`
}
