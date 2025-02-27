package sync

// Section represents a section in the Todoist application.
type Section struct {
	// ID is the ID of the section.
	ID string `json:"id"`

	// Name is the name of the section.
	Name string `json:"name"`

	// ProjectID is the ID of the project that the section resides in.
	ProjectID string `json:"project_id"`

	// SectionOrder defines the position of the section among all the sections in the project.
	SectionOrder int `json:"section_order"`

	// Collapsed indicates whether the section's tasks are collapsed.
	Collapsed bool `json:"collapsed"`

	// SyncID is a special ID for shared sections. Used internally and can be ignored.
	SyncID *string `json:"sync_id,omitempty"`

	// IsDeleted indicates whether the section is marked as deleted.
	IsDeleted bool `json:"is_deleted"`

	// IsArchived indicates whether the section is marked as archived.
	IsArchived bool `json:"is_archived"`

	// ArchivedAt is the date when the section was archived (or null if not archived).
	ArchivedAt *string `json:"archived_at,omitempty"`

	// AddedAt is the date when the section was created.
	AddedAt string `json:"added_at"`
}
