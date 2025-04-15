package sync

// Clears the locations list, which is used for location reminders.
//
// See [Clear locations] for more details.
//
// [Clear locations]: https://todoist.com/api/v1/docs#tag/Sync/Reminders/Locations
type LocationClearArgs struct{}

// Return "clear_locations" as command type.
func (args *LocationClearArgs) Type() string {
	return "clear_locations"
}
