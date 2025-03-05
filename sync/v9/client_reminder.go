package sync

func (sc *SyncClient) AddReminder(args *ReminderAddArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) UpdateReminder(args *ReminderUpdateArgs) (*Response, error) {
	return sc.executeCommand(args)
}

func (sc *SyncClient) DeleteReminder(id string) (*Response, error) {
	args := &ReminderDeleteArgs{ID: id}
	return sc.executeCommand(args)
}
