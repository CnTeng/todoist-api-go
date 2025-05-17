package sync

type ProjectMoveArgs struct {
	// Required.
	// The ID of the project (could be a temp id).
	ID string `json:"id"`

	// Optional.
	// The ID of the parent project (could be temp id). If set to null, the
	// project will be moved to the root
	ParentID *string `json:"parent_id,omitempty"`
}

func (args *ProjectMoveArgs) command() string {
	return "project_move"
}

type ProjectMoveToWorkspaceArgs struct {
	// Required.
	// The ID of the project (can be a temp id).
	ProjectID string `json:"project_id"`

	// Required.
	// The ID of the workspace the project will be moved into.
	WorkspaceID string `json:"workspace_id"`

	// Optional.
	// If true the project will be restricted access, otherwise any workspace
	// member could join it.
	IsInviteOnly *bool `json:"is_invite_only,omitempty"`

	// Optional.
	// If supplied, the project and any child projects will be moved into this
	// workspace folder
	FolderID *string `json:"folder_id,omitempty"`
}

func (args *ProjectMoveToWorkspaceArgs) command() string {
	return "project_move_to_workspace"
}

type ProjectMoveToPersonalArgs struct {
	// Required.
	// The ID of the project being moved back (can be a temp id).
	ProjectID string `json:"project_id"`
}

func (args *ProjectMoveToPersonalArgs) command() string {
	return "project_move_to_personal"
}
