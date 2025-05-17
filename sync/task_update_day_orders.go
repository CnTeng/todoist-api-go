package sync

type TaskUpdateDayOrdersArgs struct {
	// Required.
	// A dictionary, where an item id is the key, and the day_order is the value.
	IdsToOrders map[string]int `json:"ids_to_orders"`
}

func (args *TaskUpdateDayOrdersArgs) command() string {
	return "item_update_day_orders"
}
