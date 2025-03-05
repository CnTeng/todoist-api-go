package sync

// Clears the locations list, which is used for location reminders.
//
// See [Clear locations] for more details.
//
// [Clear locations]: https://developer.todoist.com/sync/v9/#clear-locations
type LoactionClearArgs struct{}

// Return "clear_locations" as command type.
func (args *LoactionClearArgs) Type() string {
	return "clear_locations"
}
