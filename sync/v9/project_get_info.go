package sync

// Get detailed information about the project, including all the notes.
//
// A full sync only returns up to 10 notes for each project. If a client
// requires more, they can be downloaded using this endpoint.
//
// It returns a JSON object with the project, and optionally the notes
// attributes. See [ProjectGetInfoResponse] for more details.
//
// See [Get project info] for more details.
//
// [Get project info]: https://developer.todoist.com/sync/v9/#get-project-info
type ProjectGetInfoParams struct {
	// Required.
	// The project's unique ID.
	ProjectID string `json:"project_id" url:"project_id"`

	// Optional.
	// Whether to return the notes of the project. Defaults to true.
	AllData *bool `json:"all_data,omitempty" url:"all_data,omitempty"`
}

// [ProjectGetInfoResponse] returns a JSON object with the project, and
// optionally the notes attributes.
type ProjectGetInfoResponse struct {
	Project *Project       `json:"project"`
	Notes   []*ProjectNote `json:"notes"`
}
