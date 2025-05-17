package sync

type FilterDeleteArgs struct {
	// Required.
	// The ID of the filter.
	ID string `json:"id"`
}

func (args *FilterDeleteArgs) command() string {
	return "filter_delete"
}
