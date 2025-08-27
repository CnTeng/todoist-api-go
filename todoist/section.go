package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

// SectionService provides methods for managing sections.
type SectionService struct {
	client *Client
}

func NewSectionService(client *Client) *SectionService {
	return &SectionService{client: client}
}

// AddSection adds a new section to a project.
//
// See [Add a section] for more details.
//
// [Add a section]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Add-a-section
func (s *SectionService) AddSection(ctx context.Context, args *sync.SectionAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// UpdateSection updates section attributes.
//
// See [Update a section] for more details.
//
// [Update a section]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Update-a-section
func (s *SectionService) UpdateSection(ctx context.Context, args *sync.SectionUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// MoveSection moves section to a different location.
//
// See [Move a section] for more details.
//
// [Move a section]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Move-a-section
func (s *SectionService) MoveSection(ctx context.Context, args *sync.SectionMoveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// ReorderSections updates section_order properties of sections in bulk.
//
// See [Reorder sections] for more details.
//
// [Reorder sections]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Reorder-sections
func (s *SectionService) ReorderSections(ctx context.Context, args *sync.SectionReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// ArchiveSection archives a section and all its child tasks.
//
// See [Archive a section] for more details.
//
// [Archive a section]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Archive-a-section
func (s *SectionService) ArchiveSection(ctx context.Context, args *sync.SectionArchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// ArchiveSections is a simple wrapper around [SectionService.ArchiveSection] that archives
// multiple.
func (s *SectionService) ArchiveSections(ctx context.Context, args []*sync.SectionArchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

// UnarchiveSection unarchives a section.
//
// See [Unarchive a section] for more details.
//
// [Unarchive a section]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Unarchive-a-section
func (s *SectionService) UnarchiveSection(ctx context.Context, args *sync.SectionUnarchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// UnarchiveSections is a simple wrapper around [SectionService.UnarchiveSection] that unarchives
// multiple.
func (s *SectionService) UnarchiveSections(ctx context.Context, args []*sync.SectionUnarchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

// DeleteSection deletes a section and all its child tasks.
//
// See [Delete a section] for more details.
//
// [Delete a section]: https://todoist.com/api/v1/docs#tag/Sync/Sections/Delete-a-section
func (s *SectionService) DeleteSection(ctx context.Context, args *sync.SectionDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// DeleteSections is a simple wrapper around [SectionService.DeleteSection] that
// deletes multiple.
func (s *SectionService) DeleteSections(ctx context.Context, args []*sync.SectionDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
