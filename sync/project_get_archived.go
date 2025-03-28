package sync

// Get the user's archived projects.
//
// See [Get archived projects] for more details.
//
// [Get archived projects]: https://developer.todoist.com/sync/v9/#get-archived-projects
type ProjectGetArchivedParams struct {
	// Optional.
	// The maximum number of archived projects to return (between 1 and 500,
	// default is 500).
	Limit *int `json:"limit,omitempty" url:"limit,omitempty"`

	// Optional.
	// The offset of the first archived project to return, for pagination purposes
	// (first page is 0).
	Offset *int `json:"offset,omitempty" url:"offset,omitempty"`
}

// [ProjectGetArchivedParams] returns the user's archived projects.
type ProjectGetArchivedResponse []*Project
