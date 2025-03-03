package sync

// Uncomplete and restore an archived item. See [Uncomplete item] for more
// details.
//
// Any ancestor items or sections will also be reinstated. Items will have the
// checked value reset.
//
// The reinstated items and sections will appear at the end of the list within
// their parent, after any previously active items.
//
// [Uncomplete item]: https://developer.todoist.com/sync/v9#uncomplete-item
type ItemUncompleteArgs struct {
	// Required.
	// Task ID to uncomplete
	ID string `json:"id"`
}

// Return "item_uncomplete" as command type.
func (args *ItemUncompleteArgs) Type() string {
	return "item_uncomplete"
}
