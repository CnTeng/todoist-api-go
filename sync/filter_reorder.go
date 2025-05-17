package sync

type FilterReorderArgs struct {
	// Required.
	// A dictionary, where a filter ID is the key, and the order its value.
	IDOrderMapping map[string]int `json:"id_order_mapping"`
}

func (args *FilterReorderArgs) command() string {
	return "filter_update_orders"
}
