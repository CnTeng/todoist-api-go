package sync

// Project represents a project in the Todoist application.
type Project struct {
	// ID is the ID of the project.
	ID string `json:"id"`

	// Name is the name of the project.
	Name string `json:"name"`

	// Color is the color of the project icon.
	Color string `json:"color"`

	// ParentID is the ID of the parent project. Set to null for root projects.
	ParentID *string `json:"parent_id,omitempty"`

	// ChildOrder defines the position of the project among all projects with the same ParentID.
	ChildOrder int `json:"child_order"`

	// Collapsed indicates whether the project's sub-projects are collapsed.
	Collapsed bool `json:"collapsed"`

	// Shared indicates whether the project is shared.
	Shared bool `json:"shared"`

	// CanAssignTasks indicates whether tasks in the project can be assigned to users.
	CanAssignTasks bool `json:"can_assign_tasks"`

	// IsDeleted indicates whether the project is marked as deleted.
	IsDeleted bool `json:"is_deleted"`

	// IsArchived indicates whether the project is marked as archived.
	IsArchived bool `json:"is_archived"`

	// IsFavorite indicates whether the project is a favorite.
	IsFavorite bool `json:"is_favorite"`

	// SyncID is the identifier to find the match between different copies of shared projects.
	SyncID *string `json:"sync_id,omitempty"`

	// InboxProject indicates whether the project is Inbox.
	InboxProject bool `json:"inbox_project,omitempty"`

	// TeamInbox indicates whether the project is TeamInbox.
	TeamInbox bool `json:"team_inbox,omitempty"`

	// ViewStyle determines the way the project is displayed (either list or board).
	ViewStyle string `json:"view_style"`
}
