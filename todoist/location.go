package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

// LocationService provides methods for managing locations.
type LocationService struct {
	client *Client
}

func NewLocationService(client *Client) *LocationService {
	return &LocationService{client: client}
}

// Clears the locations list, which is used for location reminders.
//
// See [Clear locations] for more details.
//
// [Clear locations]: https://todoist.com/api/v1/docs#tag/Sync/Reminders/Locations
func (s *LocationService) ClearLocations(ctx context.Context) (*sync.SyncResponse, error) {
	return s.client.ExecuteCommand(ctx, &sync.LocationClearArgs{})
}
