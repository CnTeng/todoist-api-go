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

// Add a filter.
//
// See [Add a filter] for more details.
//
// [Add a filter]: https://todoist.com/api/v1/docs#tag/Sync/Filters/Add-a-filter
func (s *FilterService) AddFilter(ctx context.Context, args *sync.FilterAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// Update a filter.
//
// See [Update a filter] for more details.
//
// [Update a filter]: https://todoist.com/api/v1/docs#tag/Sync/Filters/Update-a-filter
func (s *FilterService) UpdateFilter(ctx context.Context, args *sync.FilterUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// Update multiple filter orders.
//
// See [Update multiple filter orders] for more details.
//
// [Update multiple filter orders]: https://todoist.com/api/v1/docs#tag/Sync/Filters/Update-multiple-filter-orders
func (s *FilterService) ReorderFilters(ctx context.Context, args *sync.FilterReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// Delete a filter.
//
// See [Delete a filter] for more details.
//
// [Delete a filter]: https://todoist.com/api/v1/docs#tag/Sync/Filters/Delete-a-filter
func (s *FilterService) DeleteFilter(ctx context.Context, args *sync.FilterDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// A simple wrapper around [FilterService.DeleteFilter] that deletes multiple.
func (s *FilterService) DeleteFilters(ctx context.Context, args []*sync.FilterDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
