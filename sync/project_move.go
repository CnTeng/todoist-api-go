package sync

// Update parent project relationships of the project.
//
// See [Move a project] for more details.
//
// [Move a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Move-a-project
type ProjectMoveArgs struct {
	// Required.
	// The ID of the project (could be a temp id).
	ID string `json:"id"`

	// Optional.
	// The ID of the parent project (could be temp id). If set to null, the
	// project will be moved to the root
	ParentID *string `json:"parent_id,omitempty"`
}

// Return "project_move" as command type.
func (args *ProjectMoveArgs) Type() string {
	return "project_move"
}

// Moves a personal project into the target workspace.
//
//   - Moving a parent project to a workspace will also move all its child projects to that workspace.
//   - If no folder_id is supplied, child projects will be moved into a folder with the same name as the parent project being moved
//   - If a folder_id is supplied, the parent and child projects will be moved into that folder.
//   - At the moment, it is not possible to move a project to another workspace (changing its workspace_id), or to the user's personal workspace.
//   - Moving a project to a workspace affects all its collaborators. Collaborators who are not members of the target workspace will be added as guests, if guest members are allowed in the target workspace
//
// See [Move a Project to a Workspace] for more details.
//
// [Move a Project to a Workspace]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Move-a-Project-to-a-Workspace
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

// Return "project_move_to_workspace" as command type.
func (args *ProjectMoveToWorkspaceArgs) Type() string {
	return "project_move_to_workspace"
}

// Moves a project inside a workspace out back into a users personal space.
//
// Only the original creator of the workspace have permissions to do this, and
// only if they are still currently an admin of said workspace.
//
// See [Move a Project out of a Workspace] for more details.
//
// [Move a Project out of a Workspace]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Move-a-Project-out-of-a-Workspace
type ProjectMoveToPersonalArgs struct {
	// Required.
	// The ID of the project being moved back (can be a temp id).
	ProjectID string `json:"project_id"`
}

// Return "project_move_to_personal" as command type.
func (args *ProjectMoveToPersonalArgs) Type() string {
	return "project_move_to_personal"
}
