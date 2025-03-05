package sync

func (sc *SyncClient) AddSection(args *SectionAddArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) UpdateSection(args *SectionUpdateArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) MoveSection(args *SectionMoveArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) ReorderSection(args *SectionReorderArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) DeleteSection(id string) (*Response, error) {
	args := &SectionDeleteArgs{id}
	return sc.executeCommand(args)
}

func (sc *SyncClient) ArchiveSection(id string) (*Response, error) {
	args := &SectionArchiveArgs{ID: id}
	return sc.executeCommand(args)
}

func (sc *SyncClient) UnarchiveSection(id string) (*Response, error) {
	args := &SectionUnarchiveArgs{ID: id}
	return sc.executeCommand(args)
}
