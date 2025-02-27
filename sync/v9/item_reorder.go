package sync

// The command updates child_order properties of items in bulk.
type ItemReorderArgs struct {
	Items []struct {
		// The ID of the item to update
		ID string `json:"id"`

		// The new order.
		ChildOrder int `json:"child_order"`
	}
}
