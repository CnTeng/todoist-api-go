package sync

// Update multiple label orders.
//
// See [Update multiple label orders] for more details.
//
// [Update multiple label orders]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Update-multiple-label-orders
type LabelReorderArgs struct {
	// Required.
	// A dictionary, where a label id is the key, and the item_order value.
	IDOrderMapping map[string]int `json:"id_order_mapping"`
}

func (args *LabelReorderArgs) command() string {
	return "label_update_orders"
}
