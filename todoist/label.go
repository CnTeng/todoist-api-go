package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

// LabelService provides methods for managing labels.
type LabelService struct {
	client *Client
}

func NewLabelService(client *Client) *LabelService {
	return &LabelService{client: client}
}

func (s *LabelService) AddLabel(ctx context.Context, args *sync.LabelAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *LabelService) UpdateLabel(ctx context.Context, args *sync.LabelUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *LabelService) RenameSharedLabel(ctx context.Context, args *sync.LabelRenameSharedArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *LabelService) ReorderLabels(ctx context.Context, args *sync.LabelReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *LabelService) DeleteLabel(ctx context.Context, args *sync.LabelDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *LabelService) DeleteLabels(ctx context.Context, args []*sync.LabelDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

func (s *LabelService) DeleteSharedLabel(ctx context.Context, args *sync.LabelDeleteSharedArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *LabelService) DeleteSharedLabels(ctx context.Context, args []*sync.LabelDeleteSharedArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
