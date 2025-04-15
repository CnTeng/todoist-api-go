package sync

// Update multiple filter orders.
//
// See [Update multiple filter orders] for more details.
//
// [Update multiple filter orders]: https://todoist.com/api/v1/docs#tag/Sync/Filters/Update-multiple-filter-orders
type FilterReorderArgs struct {
	// Required.
	// A dictionary, where a filter ID is the key, and the order its value.
	IDOrderMapping map[string]int `json:"id_order_mapping"`
}

// Return "filter_update_orders" as command type.
func (args *FilterReorderArgs) Type() string {
	return "filter_update_orders"
}
