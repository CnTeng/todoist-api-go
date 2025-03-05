package sync

func (sc *SyncClient) AddLabel(args *LabelAddArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) UpdateLabel(args *LabelUpdateArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) DeleteLabel(id string) (*Response, error) {
	args := &LabelDeleteArgs{ID: id}
	return sc.executeCommand(args)
}

func (sc *SyncClient) RenameSharedLabel(args *LabelRenameSharedArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) DeleteSharedLabel(name string) (*Response, error) {
	args := &LabelDeleteSharedArgs{Name: name}
	return sc.executeCommand(args)
}

func (sc *SyncClient) ReorderLabels(args *LabelReorderArgs) (*Response, error) {
	return sc.executeCommand(args)
}
