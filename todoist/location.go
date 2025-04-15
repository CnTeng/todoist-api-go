package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (sc *Client) ClearLocations(ctx context.Context) (*sync.SyncResponse, error) {
	args := &sync.LoactionClearArgs{}
	return sc.ExecuteCommand(ctx, args)
}
