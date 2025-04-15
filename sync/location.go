package sync

// Locations are a top-level entity in the sync model. They contain a list of
// all locations that are used within user's current location reminders.
//
// The location object is specific, as it's not an object, but an ordered array.
//
//   - 0: Name of the location.
//   - 1: Location latitude.
//   - 2: Location longitude.
//
// See [Locations] for more details.
//
// [locations]: https://todoist.com/api/v1/docs#tag/Sync/Reminders/Locations
type Location []string
