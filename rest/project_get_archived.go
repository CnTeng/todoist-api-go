package rest

import "github.com/CnTeng/todoist-api-go/sync"

type ProjectGetArchivedParams struct {
	// Optional.
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	// Optional.
	// Default is 50. The number of objects to return in a page.
	Limit *int `json:"limit,omitempty" url:"limit,omitempty"`
}

type ProjectGetArchivedResponse struct {
	Projects   []*sync.Project `json:"results"`
	NextCursor *string         `json:"next_cursor,omitempty"`
}
