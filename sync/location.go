package sync

// Contain a list of all locations that are used within user's current location
// reminders:
//   - 0: Name of the location.
//   - 1: Location latitude.
//   - 2: Location longitude.
//
// See [Locations] for more details.
//
// [locations]: https://developer.todoist.com/sync/v9/#locations
type Location []string
