package sync

// Update the day orders of multiple tasks at once.
//
// See [Update day orders] for more details.
//
// [Update day orders]: https://todoist.com/api/v1/docs#tag/Sync/Tasks/Update-day-orders
type TaskUpdateDayOrdersArgs struct {
	// Required.
	// A dictionary, where an item id is the key, and the day_order is the value.
	IdsToOrders map[string]int `json:"ids_to_orders"`
}

// Return "item_update_day_orders" as command type.
func (args *TaskUpdateDayOrdersArgs) Type() string {
	return "item_update_day_orders"
}
