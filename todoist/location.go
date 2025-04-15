package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (sc *Client) ClearLocations(ctx context.Context) (*sync.SyncResponse, error) {
	return sc.ExecuteCommand(ctx, &sync.LocationClearArgs{})
}
