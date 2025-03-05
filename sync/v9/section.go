package sync

// Section represents a section.
//
// See [Sections] for more details.
//
// [Sections]: https://developer.todoist.com/sync/v9/#sections
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
	Collapsed bool `json:"collapsed"`

	// A special ID for shared sections (a number or null if not set). Used
	// internally and can be ignored.
	SyncID *string `json:"sync_id"`

	// Whether the section is marked as deleted
	IsDeleted bool `json:"is_deleted"`

	// Whether the section is marked as archived.
	IsArchived bool `json:"is_archived"`

	// The date when the section was archived.
	ArchivedAt *string `json:"archived_at"`

	// The date when the section was created.
	AddedAt string `json:"added_at"`
}
