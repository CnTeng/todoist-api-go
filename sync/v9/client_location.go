package sync

func (sc *SyncClient) ClearLocations() (*Response, error) {
	args := &LoactionClearArgs{}
	return sc.executeCommand(args)
}
