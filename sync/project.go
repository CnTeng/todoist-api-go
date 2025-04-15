package sync

import "time"

// Project represents a project.
//
// See [Projects] for more details.
//
// [Projects]: https://todoist.com/api/v1/docs#tag/Sync/Projects
type Project struct {
	// The ID of the project.
	ID string `json:"id"`

	// The name of the project.
	Name string `json:"name"`

	// Description for the project. Only used for teams.
	Description string `json:"description"`

	// Real or temp ID of the workspace the project. Only used for teams.
	WorkspaceID string `json:"workspace_id"`

	// Indicates if the project is invite-only or if it should be visible for
	// everyone in the workspace. If missing or null, the default value from the
	// workspace is_invite_only_default will be used. Only used for teams.
	IsInviteOnly bool `json:"is_invite_only"`

	// The status of the project. Possible values: PLANNED, IN_PROGRESS, PAUSED,
	// COMPLETED, CANCELED. Only used for teams.
	Status string `json:"status"`

	// If False, the project is invite-only and people can't join by link. If
	// true, the project is visible to anyone with a link, and anyone can join it.
	// Only used for teams.
	IsLinkSharingEnabled bool `json:"is_link_sharing_enabled"`

	// The role a user can have. Possible values: CREATOR, ADMIN, READ_WRITE,
	// READ_AND_COMMENT, READ_ONLY. (CREATOR is equivalent to admin and is
	// automatically set at project creation; it can’t be set to anyone later).
	// Defaults to READ_WRITE. Only used for teams.
	CollaboratorRoleDefault string `json:"collaborator_role_default"`

	// The color of the project icon. Refer to the name column in the
	// https://todoist.com/api/v1/docs#tag/Colors guide for more info.
	Color Color `json:"color"`

	// The ID of the parent project. Set to null for root projects.
	ParentID *string `json:"parent_id,omitempty"`

	// The order of the project. Defines the position of the project among all the
	// projects with the same parent_id.
	ChildOrder int `json:"child_order"`

	// Whether the project's sub-projects are collapsed (a true or false value).
	Collapsed bool `json:"collapsed"`

	// Whether the project is shared (a true or false value).
	Shared bool `json:"shared"`

	// Whether tasks in the project can be assigned to users (a true or false
	// value).
	CanAssignTasks bool `json:"can_assign_tasks"`

	// Whether the project is marked as deleted (a true or false value).
	IsDeleted bool `json:"is_deleted"`

	// Whether the project is marked as archived (a true or false value).
	IsArchived bool `json:"is_archived"`

	// Whether the project is a favorite (a true or false value).
	IsFavorite bool `json:"is_favorite"`

	// Workspace or personal projects from a cancelled subscription (a true or
	// false value).
	IsFrozen bool `json:"is_frozen"`

	// The mode in which to render tasks in this project. One of list, board, or
	// calendar.
	ViewStyle string `json:"view_style"`

	// The role of the requesting user. Possible values: CREATOR, ADMIN,
	// READ_WRITE, READ_AND_COMMENT, READ_ONLY. Only used for teams
	Role string `json:"role"`

	// Whether the project is Inbox (true or otherwise this property is not sent).
	InboxProject bool `json:"inbox_project,omitempty"`

	// The ID of the folder which this project is in.
	FolderId *string `json:"folder_id,omitempty"`

	// Project creation date in the format YYYY-MM-DDTHH:MM:SSZ date.
	CreatedAt time.Time `json:"created_at"`

	// Project update date in the format YYYY-MM-DDTHH:MM:SSZ date.
	UpdatedAt time.Time `json:"updated_at"`
}
