package sync

// Update the day orders of multiple items at once.
type ItemUpdateDayOrdersArgs struct {
	// A dictionary, where an item id is the key, and the day_order value:
	//   item_id: day_order.
	IdsToOrders map[string]int `json:"ids_to_orders"` // required
}
