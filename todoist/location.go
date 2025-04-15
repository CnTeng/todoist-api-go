package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) ClearLocations(ctx context.Context) (*sync.SyncResponse, error) {
	return c.ExecuteCommand(ctx, &sync.LocationClearArgs{})
}
