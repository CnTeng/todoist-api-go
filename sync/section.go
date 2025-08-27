package sync

import "time"

// Section represents a section in Todoist.
//
// See [Sections] for more details.
//
// [Sections]: https://todoist.com/api/v1/docs#tag/Sync/Sections
type Section struct {
	// The ID of the section.
	ID string `json:"id"`

	// The name of the section.
	Name string `json:"name"`

	// Project that the section resides in.
	ProjectID string `json:"project_id"`

	// The order of the section. Defines the position of the section among all the
	// sections in the project.
	SectionOrder int `json:"section_order"`

	// Whether the section's tasks are collapsed.
	IsCollapsed bool `json:"is_collapsed"`

	// A special ID for shared sections (a number or null if not set). Used
	// internally and can be ignored.
	SyncID *string `json:"sync_id"`

	// Whether the section is marked as deleted (a true or false value).
	IsDeleted bool `json:"is_deleted"`

	// Whether the section is marked as archived (a true or false value).
	IsArchived bool `json:"is_archived"`

	// The date when the section was archived (or null if not archived).
	ArchivedAt *time.Time `json:"archived_at"`

	// The date when the section was created.
	AddedAt time.Time `json:"added_at"`

	// The date when the section was updated.
	UpdatedAt time.Time `json:"updated_at"`
}
