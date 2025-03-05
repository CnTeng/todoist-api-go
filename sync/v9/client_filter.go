package sync

func (sc *SyncClient) AddFilter(args *FilterAddArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) UpdateFilter(args *FilterUpdateArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) DeleteFilter(id string) (*Response, error) {
	args := &FilterDeleteArgs{ID: id}
	return sc.executeCommand(args)
}

func (sc *SyncClient) ReorderFilters(args *FilterReorderArgs) (*Response, error) {
	return sc.executeCommand(args)
}
