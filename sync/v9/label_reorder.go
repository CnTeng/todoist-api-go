package sync

// Update multiple label orders.
//
// See [Update multiple label orders] for more details.
//
// [Update multiple label orders]: https://developer.todoist.com/sync/v9/#update-multiple-label-orders
type LabelReorderArgs struct {
	// Required.
	// A dictionary, where a label id is the key, and the item_order value.
	IDOrderMapping map[string]int `json:"id_order_mapping"`
}

// Return "label_update_orders" as command type.
func (args *LabelReorderArgs) Type() string {
	return "label_update_orders"
}
