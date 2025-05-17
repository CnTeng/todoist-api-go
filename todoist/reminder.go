package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

// ReminderService provides methods for managing reminders.
type ReminderService struct {
	client *Client
}

func NewReminderService(client *Client) *ReminderService {
	return &ReminderService{client: client}
}

// Add a new reminder to the user account related to the API credentials.
//
// See [Add a reminder] for more details.
//
// [Add a reminder]: https://todoist.com/api/v1/docs#tag/Sync/Reminders/Add-a-reminder
func (s *ReminderService) AddReminder(ctx context.Context, args *sync.ReminderAddArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// Update a reminder from the user account related to the API credentials.
//
// See [Update a reminder] for more details.
//
// [Update a reminder]: https://todoist.com/api/v1/docs#tag/Sync/Reminders/Update-a-reminder
func (s *ReminderService) UpdateReminder(ctx context.Context, args *sync.ReminderUpdateArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// Delete a reminder from the current user account.
//
// See [Delete a reminder] for more details.
//
// [Delete a reminder]: https://todoist.com/api/v1/docs#tag/Sync/Reminders/Delete-a-reminder
func (s *ReminderService) DeleteReminder(ctx context.Context, args *sync.ReminderDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, args)
}

// A simple wrapper around [ReminderService.DeleteReminder] that deletes
// multiple.
func (s *ReminderService) DeleteReminders(ctx context.Context, args []*sync.ReminderDeleteArgs) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommands(ctx, sync.NewCommands(args))
}
