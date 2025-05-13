package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

// FilterService provides methods for managing filters.
type FilterService struct {
	client *Client
}

func NewFilterService(client *Client) *FilterService {
	return &FilterService{client: client}
}

func (s *FilterService) AddFilter(ctx context.Context, args *sync.FilterAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *FilterService) UpdateFilter(ctx context.Context, args *sync.FilterUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *FilterService) ReorderFilters(ctx context.Context, args *sync.FilterReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *FilterService) DeleteFilter(ctx context.Context, args *sync.FilterDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

func (s *FilterService) DeleteFilters(ctx context.Context, args []*sync.FilterDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
