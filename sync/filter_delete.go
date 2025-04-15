package sync

// Delete a filter.
//
// See [Delete a filter] for more details.
//
// [Delete a filter]: https://todoist.com/api/v1/docs#tag/Sync/Filters/Delete-a-filter
type FilterDeleteArgs struct {
	// Required.
	// The ID of the filter.
	ID string `json:"id"`
}

// Return "filter_delete" as command type.
func (args *FilterDeleteArgs) Type() string {
	return "filter_delete"
}
