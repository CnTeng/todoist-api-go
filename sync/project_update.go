package sync

// Update an existing project. See [Update a project] for more details.
//
// [Update a project]: https://developer.todoist.com/sync/v9/#update-a-project
type ProjectUpdateArgs struct {
	// Required.
	// The ID of the project.
	ID string `json:"id"`

	// Optional.
	// The name of the project
	Name *string `json:"name,omitempty"`

	// Optional.
	// The color of the project icon.
	//
	// See https://developer.todoist.com/guides/#colors for a list of available.
	Color *Color `json:"color,omitempty"`

	// Optional.
	// Whether the project's sub-projects are collapsed.
	Collapsed *bool `json:"collapsed,omitempty"`

	// Optional.
	// Whether the project is a favorite.
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// Optional.
	// A string value (either list or board). This determines the way the project
	// is displayed within the Todoist clients.
	ViewStyle *string `json:"view_style,omitempty"`
}

// Return "project_update" as command type.
func (args *ProjectUpdateArgs) Type() string {
	return "project_update"
}
