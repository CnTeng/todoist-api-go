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

func (s *SectionService) AddSection(ctx context.Context, args *sync.SectionAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *SectionService) UpdateSection(ctx context.Context, args *sync.SectionUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *SectionService) MoveSection(ctx context.Context, args *sync.SectionMoveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *SectionService) ReorderSections(ctx context.Context, args *sync.SectionReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *SectionService) ArchiveSection(ctx context.Context, args *sync.SectionArchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *SectionService) ArchiveSections(ctx context.Context, args []*sync.SectionArchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (s *SectionService) UnarchiveSection(ctx context.Context, args *sync.SectionUnarchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *SectionService) UnarchiveSections(ctx context.Context, args []*sync.SectionUnarchiveArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (s *SectionService) DeleteSection(ctx context.Context, args *sync.SectionDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *SectionService) DeleteSections(ctx context.Context, args []*sync.SectionDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
