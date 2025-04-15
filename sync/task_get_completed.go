package sync

import "time"

// Retrieves a list of completed tasks strictly limited by the specified
// completion date range (up to 3 months).
//
// It can retrieve completed items:
//
//   - From all the projects the user has joined in a workspace
//   - From all the projects of the user
//   - That match many [supported filters]
//
// By default, the response is limited to a page containing a maximum of 50
// items (configurable using limit).
//
// Subsequent pages of results can be fetched by using the next_cursor value
// from the response as the cursor value for the next request.
//
// See [Tasks Completed By Completion Date] for more details.
//
// [supported filters]: https://todoist.com/help/articles/introduction-to-filters-V98wIH
// [Tasks Completed By Completion Date]: https://todoist.com/api/v1/docs#tag/Tasks/operation/tasks_completed_by_completion_date_api_v1_tasks_completed_by_completion_date_get
type TaskGetCompletedByCompletionDateParams struct {
	// Required.
	Since time.Time `json:"since" url:"since"`

	// Required.
	Until time.Time `json:"until" url:"until"`

	// Optional.
	WorkspaceID *string `json:"workspace_id,omitempty" url:"workspace_id,omitempty"`

	// Optional.
	FilterQuery *string `json:"filter_query,omitempty" url:"filter_query,omitempty"`

	// Optional.
	FilterLang *string `json:"filter_lang,omitempty" url:"filter_lang,omitempty"`

	// Optional.
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	// Optional.
	// Default is 50.
	Limit *int `json:"limit,omitempty" url:"limit,omitempty"`
}

// Retrieves a list of completed items strictly limited by the specified due
// date range (up to 6 weeks).
//
// It can retrieve completed items:
//
//   - From within a project, section, or parent item
//   - From all the projects the user has joined in a workspace
//   - From all the projects of the user
//   - That match many [supported filters]
//
// By default, the response is limited to a page containing a maximum of 50
// items (configurable using limit).
//
// Subsequent pages of results can be fetched by using the next_cursor value
// from the response as the cursor value for the next request.
//
// See [Tasks Completed By Due Date] for more details.
//
// [supported filters]: https://todoist.com/help/articles/introduction-to-filters-V98wIH
// [Tasks Completed By Due Date]: https://todoist.com/api/v1/docs#tag/Tasks/operation/tasks_completed_by_due_date_api_v1_tasks_completed_by_due_date_get
type TaskGetCompletedByDueDateParams struct {
	// Required.
	Since time.Time `json:"since" url:"since"`

	// Required.
	Until time.Time `json:"until" url:"until"`

	// Optional.
	WorkspaceID *int `json:"workspace_id,omitempty" url:"workspace_id,omitempty"`

	// Optional.
	ProjectID *string `json:"project_id,omitempty" url:"project_id,omitempty"`

	// Optional.
	SectionID *string `json:"section_id,omitempty" url:"section_id,omitempty"`

	// Optional.
	ParentID *string `json:"parent_id,omitempty" url:"parent_id,omitempty"`

	// Optional.
	FilterQuery *string `json:"filter_query,omitempty" url:"filter_query,omitempty"`

	// Optional.
	FilterLang *string `json:"filter_lang,omitempty" url:"filter_lang,omitempty"`

	// Optional.
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	// Optional.
	// Default is 50.
	Limit *int `json:"limit,omitempty" url:"limit,omitempty"`
}

type TaskGetCompletedResponse struct {
	Tasks      []*Task `json:"items"`
	NextCursor *string `json:"next_cursor,omitempty"`
}
