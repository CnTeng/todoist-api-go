package sync

// Update an existing project.
//
// See [Update a project] for more details.
//
// [Update a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Update-a-project
type ProjectUpdateArgs struct {
	// Required.
	// The ID of the project (could be temp id).
	ID string `json:"id"`

	// Optional.
	// The name of the project (a string value).
	Name *string `json:"name,omitempty"`

	// Optional.
	// The color of the project icon. Refer to the name column in the
	// https://todoist.com/api/v1/docs#tag/Colors guide for more info.
	Color *Color `json:"color,omitempty"`

	// Optional.
	// Whether the project's sub-projects are collapsed (a true or false value).
	Collapsed *bool `json:"collapsed,omitempty"`

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
	// Defaults to READ_WRITE. Only used for teams.
	CollaboratorRoleDefault *string `json:"collaborator_role_default,omitempty"`
}

// Return "project_update" as command type.
func (args *ProjectUpdateArgs) Type() string {
	return "project_update"
}
