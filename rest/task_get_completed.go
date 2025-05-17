package rest

import (
	"time"

	"github.com/CnTeng/todoist-api-go/sync"
)

type TaskGetCompletedByCompletionDateParams struct {
	// Required.
	Since time.Time `json:"since" url:"since"`

	// Required.
	Until time.Time `json:"until" url:"until"`

	// Optional.
	WorkspaceID *string `json:"workspace_id,omitempty" url:"workspace_id,omitempty"`

	// Optional.
	ProjectID *string `json:"project_id,omitempty" url:"project_id,omitempty"`

	// Optional.
	SectionID *string `json:"section_id,omitempty" url:"section_id,omitempty"`

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
	Tasks      []*sync.Task `json:"items"`
	NextCursor *string      `json:"next_cursor,omitempty"`
}
