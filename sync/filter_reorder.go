package sync

// Update the orders of multiple filters at once.
//
// See [Update multiple filter orders] for more details.
//
// [Update multiple filter orders]: https://developer.todoist.com/sync/v9/#update-multiple-filter-orders
type FilterReorderArgs struct {
	// Required.
	// A dictionary, where a filter ID is the key, and the order its value.
	IDOrderMapping map[string]int `json:"id_order_mapping"`
}

// Return "filter_update_orders" as command type.
func (args *FilterReorderArgs) Type() string {
	return "filter_update_orders"
}
