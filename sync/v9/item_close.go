package sync

// A simplified version of item_complete / item_update_date_complete.
// The command does exactly what official clients do when you close a task:
//
//	regular tasks are completed and moved to the archive
//	recurring tasks are scheduled to their next occurrence
type ItemCloseArgs struct {
	// The ID of the item to close (a number or a temp id).
	ID string `json:"id"` // required
}

func NewItemCloseCommand(args ItemCloseArgs) *Command {
	return NewCommand("item_close", args)
}
