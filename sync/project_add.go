package sync

// Add a new project. See [Add a project] for more details.
//
// [Add a project]: https://developer.todoist.com/sync/v9/#add-a-project
type ProjectAddArgs struct {
	// Required.
	// The name of the project
	Name string `json:"name"`

	// Optional.
	// The color of the project icon.
	//
	// See https://developer.todoist.com/guides/#colors for a list of available.
	Color *Color `json:"color,omitempty"`

	// Optional.
	// The ID of the parent project. Set to null for root projects.
	ParentID *string `json:"parent_id,omitempty"`

	// Optional.
	// The order of the project. Defines the position of the project among all the
	// projects with the same parent_id.
	ChildOrder *int `json:"child_order,omitempty"`

	// Optional.
	// Whether the project is a favorite.
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// Optional.
	// A string value (either list or board). This determines the way the project
	// is displayed within the Todoist clients.
	ViewStyle *string `json:"view_style,omitempty"`
}

// Return "project_add" as command type.
func (args *ProjectAddArgs) Type() string {
	return "project_add"
}
