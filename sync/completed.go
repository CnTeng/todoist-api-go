package sync

// Get all the user's completed items (tasks). See [Get all completed items] for
// more details.
//
// [Get all completed items]: https://developer.todoist.com/sync/v9#get-all-completed-items
type CompletedGetParams struct {
	// Optional.
	// Filter the tasks by project ID.
	ProjectID *string `json:"project_id,omitempty" url:"project_id,omitempty"`

	// Optional.
	// The number of items to return (where the default is 30, and the maximum is
	// 200).
	Limit *int `json:"limit,omitempty" url:"limit,omitempty"`

	// Optional.
	// Can be used for pagination, when more than the limit number of tasks are
	// returned.
	Offset *int `json:"offset,omitempty" url:"offset,omitempty"`

	// Optional.
	// Return items with a completed date same or older than until (a string value
	// formatted as 2021-4-29T10:13:00).
	Until *string `json:"until,omitempty" url:"until,omitempty"`

	// Optional.
	// Return items with a completed date newer than since (a string value
	// formatted as 2021-4-29T10:13:00).
	Since *string `json:"since,omitempty" url:"since,omitempty"`

	// Optional.
	// Return notes together with the completed items.
	AnnotateNotes *bool `json:"annotate_notes,omitempty" url:"annotate_notes,omitempty"`

	// Optional.
	// Returns the full user item within the payload.
	AnnotateItems *bool `json:"annotate_items,omitempty" url:"annotate_items,omitempty"`
}

// [CompletedGetParams] returns a JSON object with the completed items, projects
// and sections.
type CompletedGetResponse struct {
	Items    []*CompletedTask    `json:"items"`
	Projects map[string]*Project `json:"projects"`
	Sections map[string]*Section `json:"sections"`
}
