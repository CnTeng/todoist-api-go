package sync

// Get a JSON object with the project, its notes, sections and any uncompleted
// items. See [ProjectGetDataResponse] for more details.
//
// See [Get project data] for more details.
//
// [Get project data]: https://developer.todoist.com/sync/v9/#get-project-info
type ProjectGetDataParams struct {
	// Required.
	// The project's unique ID.
	ProjectID string `json:"project_id" url:"project_id"`
}

// [ProjectGetDataParams] returns a JSON object with the project, its notes,
// sections and any uncompleted items.
type ProjectGetDataResponse struct {
	Project      *Project       `json:"project"`
	Items        []*Task        `json:"items"`
	Sections     []*Section     `json:"sections"`
	ProjectNotes []*ProjectNote `json:"project_notes"`
}
