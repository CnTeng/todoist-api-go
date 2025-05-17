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

// Add a new personal label.
//
// See [Add a personal label] for more details.
//
// [Add a personal label]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Add-a-personal-label
func (s *LabelService) AddLabel(ctx context.Context, args *sync.LabelAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// Update a personal label.
//
// See [Update a personal label] for more details.
//
// [Update a personal label]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Update-a-personal-label
func (s *LabelService) UpdateLabel(ctx context.Context, args *sync.LabelUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// Rename a shared label.
//
// See [Rename a shared label] for more details.
//
// [Rename a shared label]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Rename-a-shared-label
func (s *LabelService) RenameSharedLabel(ctx context.Context, args *sync.LabelRenameSharedArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// Update multiple label orders.
//
// See [Update multiple label orders] for more details.
//
// [Update multiple label orders]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Update-multiple-label-orders
func (s *LabelService) ReorderLabels(ctx context.Context, args *sync.LabelReorderArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// Delete a personal label.
//
// See [Delete a personal label] for more details.
//
// [Delete a personal label]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Delete-a-personal-label
func (s *LabelService) DeleteLabel(ctx context.Context, args *sync.LabelDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// A simple wrapper around [LabelService.DeleteLabel] that deletes multiple.
func (s *LabelService) DeleteLabels(ctx context.Context, args []*sync.LabelDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}

// Deletes all occurrences of a shared label from any active tasks.
//
// See [Delete shared label occurrences] for more details.
//
// [Delete shared label occurrences]: https://todoist.com/api/v1/docs#tag/Sync/Labels/Delete-shared-label-occurrences
func (s *LabelService) DeleteSharedLabel(ctx context.Context, args *sync.LabelDeleteSharedArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// A simple wrapper around [LabelService.DeleteSharedLabel] that deletes
// multiple.
func (s *LabelService) DeleteSharedLabels(ctx context.Context, args []*sync.LabelDeleteSharedArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
