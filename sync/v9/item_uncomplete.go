package sync

// This command is used to uncomplete and restore an archived item.
//
// Any ancestor items or sections will also be reinstated.
// Items will have the checked value reset.
//
// The reinstated items and sections will appear at the end of the list within their parent,
// after any previously active items.
type ItemUncompleteArgs struct {
	// Task ID to uncomplete
	ID string `json:"id"` // required
}
