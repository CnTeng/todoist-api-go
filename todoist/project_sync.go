package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

// ProjectService provides methods to manage projects in Todoist.
type ProjectService struct {
	client *Client
}

func NewProjectService(client *Client) *ProjectService {
	return &ProjectService{client: client}
}

// AddProject adds a new project.
//
// See [Add a project] for more details.
//
// [Add a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Add-a-project
func (s *ProjectService) AddProject(ctx context.Context, args *sync.ProjectAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// UpdateProject updates a project.
//
// See [Update a project] for more details.
//
// [Update a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Update-a-project
func (s *ProjectService) UpdateProject(ctx context.Context, args *sync.ProjectUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// MoveProject updates parent project relationships of the project.
//
// See [Move a project] for more details.
//
// [Move a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Move-a-project
func (s *ProjectService) MoveProject(ctx context.Context, args *sync.ProjectMoveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// MoveProjectToWorkspace moves a personal project into the target workspace.
//
//   - Moving a parent project to a workspace will also move all its child
//     projects to that workspace.
//   - If no folder_id is supplied, child projects will be moved into a folder
//     with the same name as the parent project being moved
//   - If a folder_id is supplied, the parent and child projects will be moved
//     into that folder.
//   - At the moment, it is not possible to move a project to another workspace
//     (changing its workspace_id), or to the user's personal workspace.
//   - Moving a project to a workspace affects all its collaborators.
//     Collaborators who are not members of the target workspace will be added
//     as guests, if guest members are allowed in the target workspace
//
// See [Move a Project to a Workspace] for more details.
//
// [Move a Project to a Workspace]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Move-a-Project-to-a-Workspace
func (s *ProjectService) MoveProjectToWorkspace(ctx context.Context, args *sync.ProjectMoveToWorkspaceArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// MoveProjectToPersonal moves a project inside a workspace back into a
// user's personal space.
//
// Only the original creator of the workspace has permissions to do this, and
// only if they are still currently an admin of said workspace.
//
// See [Move a Project out of a Workspace] for more details.
//
// [Move a Project out of a Workspace]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Move-a-Project-out-of-a-Workspace
func (s *ProjectService) MoveProjectToPersonal(ctx context.Context, args *sync.ProjectMoveToPersonalArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// LeaveProject is used to remove self from a workspace project. As this is a
// v2-only command, it will only accept v2 project id.
//
// See [Leave a project] for more details.
//
// [Leave a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Leave-a-project
func (s *ProjectService) LeaveProject(ctx context.Context, args *sync.ProjectLeaveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// ReorderProjects updates child_order properties of projects in bulk.
//
// See [Reorder projects] for more details.
//
// [Reorder projects]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Reorder-projects
func (s *ProjectService) ReorderProjects(ctx context.Context, args *sync.ProjectReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// ArchiveProject archives a project and its descendants.
//
// See [Archive a project] for more details.
//
// [Archive a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Archive-a-project
func (s *ProjectService) ArchiveProject(ctx context.Context, args *sync.ProjectArchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// ArchiveProjects is a simple wrapper around [ProjectService.ArchiveProject]
// that archives multiple.
func (s *ProjectService) ArchiveProjects(ctx context.Context, args []*sync.ProjectArchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

// UnarchiveProject unarchives a project. No ancestors will be unarchived along
// with the unarchived project. Instead, the project is unarchived alone, loses
// any parent relationship (becomes a root project), and is placed at the end of
// the list of other root projects.
//
// See [Unarchive a project] for more details.
//
// [Unarchive a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Unarchive-a-project
func (s *ProjectService) UnarchiveProject(ctx context.Context, args *sync.ProjectUnarchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// UnarchiveProjects is a simple wrapper around
// [ProjectService.UnarchiveProject] that unarchives multiple.
func (s *ProjectService) UnarchiveProjects(ctx context.Context, args []*sync.ProjectUnarchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

// DeleteProject deletes an existing project and all its descendants. Workspace
// projects can only be deleted by ADMINs and it must be archived first.
//
// See [Delete a project] for more details.
//
// [Delete a project]: https://todoist.com/api/v1/docs#tag/Sync/Projects/Delete-a-project
func (s *ProjectService) DeleteProject(ctx context.Context, args *sync.ProjectDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// DeleteProjects is a simple wrapper around [ProjectService.DeleteProject] that
// deletes multiple.
func (s *ProjectService) DeleteProjects(ctx context.Context, args []*sync.ProjectDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
