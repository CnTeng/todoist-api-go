package sync

// Update the day orders of multiple items at once. See [Update day orders] for
// more details.
//
// [Update day orders]: https://developer.todoist.com/sync/v9#update-day-orders
type ItemUpdateDayOrdersArgs struct {
	// Required.
	// A dictionary, where an item id is the key, and the day_order is the value.
	IdsToOrders map[string]int `json:"ids_to_orders"`
}

// Return "item_update_day_orders" as command type.
func (args *ItemUpdateDayOrdersArgs) Type() string {
	return "item_update_day_orders"
}
