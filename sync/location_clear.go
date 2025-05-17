package sync

type LocationClearArgs struct{}

func (args *LocationClearArgs) command() string {
	return "clear_locations"
}
