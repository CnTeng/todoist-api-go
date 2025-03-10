package sync

import "context"

func (sc *Client) ClearLocations(ctx context.Context) (*SyncResponse, error) {
	args := &LoactionClearArgs{}
	return sc.ExecuteCommand(ctx, args)
}
