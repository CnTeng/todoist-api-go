package sync

// Add a new project.
//
// See [Add a project] for more details.
//
// [Add a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Add-a-project
type ProjectAddArgs struct {
	// Required.
	// The name of the project (a string value).
	Name string `json:"name"`

	// Optional.
	// The color of the project icon. Refer to the name column in the
	// https://todoist.com/api/v1/docs#tag/Colors guide for more info.
	Color *Color `json:"color,omitempty"`

	// Optional.
	// The ID of the parent project. Set to null for root projects.
	ParentID *string `json:"parent_id,omitempty"`

	// Optional.
	// The ID of the folder, when creating projects in workspaces. Set to null for
	// root projects.
	FolderId *string `json:"folder_id,omitempty"`

	// Optional.
	// The order of the project. Defines the position of the project among all the
	// projects with the same parent_id.
	ChildOrder *int `json:"child_order,omitempty"`

	// Optional.
	// Whether the project is a favorite (a true or false value).
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// Optional.
	// The mode in which to render tasks in this project. One of list, board, or
	// calendar.
	ViewStyle *string `json:"view_style,omitempty"`

	// Optional.
	// Description for the project (up to 1024 characters). Only used for teams.
	Description *string `json:"description,omitempty"`

	// Optional.
	// Real or temp ID of the workspace the project. Only used for teams.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	// Optional.
	// Indicates if the project is invite-only or if it should be visible for
	// everyone in the workspace. If missing or null, the default value from the
	// workspace is_invite_only_default will be used. Only used for teams.
	IsInviteOnly *bool `json:"is_invite_only,omitempty"`

	// Optional.
	// The status of the project. Possible values: PLANNED, IN_PROGRESS, PAUSED,
	// COMPLETED, CANCELED. Only used for teams.
	Status *string `json:"status,omitempty"`

	// Optional.
	// If False, the project is invite-only and people can't join by link. If
	// true, the project is visible to anyone with a link, and anyone can join it.
	// Only used for teams.
	IsLinkSharingEnabled *bool `json:"is_link_sharing_enabled,omitempty"`

	// Optional.
	// The role a user can have. Possible values: CREATOR, ADMIN, READ_WRITE,
	// READ_AND_COMMENT, READ_ONLY. (CREATOR is equivalent to admin and is
	// automatically set at project creation; it can’t be set to anyone later).
	// Only used for teams.
	CollaboratorRoleDefault *string `json:"collaborator_role_default,omitempty"`
}

// Return "project_add" as command type.
func (args *ProjectAddArgs) Type() string {
	return "project_add"
}
