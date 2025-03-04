package sync

// Project represents a project. See [Projects] for more details.
//
// [Projects]: https://developer.todoist.com/sync/v9/#projects
type Project struct {
	// The ID of the project.
	ID string `json:"id"`

	// The name of the project.
	Name string `json:"name"`

	// The color of the project icon.
	//
	// See https://developer.todoist.com/guides/#colors for a list of available.
	Color Color `json:"color"`

	// The ID of the parent project. Set to null for root projects.
	ParentID *string `json:"parent_id,omitempty"`

	// The order of the project. Defines the position of the project among all the
	// projects with the same parent_id.
	ChildOrder int `json:"child_order"`

	// Whether the project's sub-projects are collapsed.
	Collapsed bool `json:"collapsed"`

	// Whether the project is shared.
	Shared bool `json:"shared"`

	// Whether tasks in the project can be assigned to users.
	CanAssignTasks bool `json:"can_assign_tasks"`

	// Whether the project is marked as deleted.
	IsDeleted bool `json:"is_deleted"`

	// Whether the project is marked as archived.
	IsArchived bool `json:"is_archived"`

	// Whether the project is a favorite.
	IsFavorite bool `json:"is_favorite"`

	// Identifier to find the match between different copies of shared projects.
	// When you share a project, its copy has a different ID for your
	// collaborators. To find a project in a different account that matches yours,
	// you can use the "sync_id" attribute. For non-shared projects the attribute
	// is set to null.
	SyncID *string `json:"sync_id,omitempty"`

	// Whether the project is Inbox.
	InboxProject bool `json:"inbox_project,omitempty"`

	// Whether the project is TeamInbox.
	TeamInbox bool `json:"team_inbox,omitempty"`

	// A string value (either list or board). This determines the way the project
	// is displayed within the Todoist clients.
	ViewStyle string `json:"view_style"`
}
