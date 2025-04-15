package sync

// There are two types of labels that can be added to Todoist tasks. We refer to
// these as personal and shared labels.
//
// # Personal labels
//
// Labels created by the current user will show up in their personal label list.
// These labels can be customized and will stay in their account unless deleted.
//
// A personal label can be converted to a shared label by the user if they no
// longer require them to be stored against their account, but they still appear
// on shared tasks.
//
// # Shared labels
//
// A label created by a collaborator that doesn't share a name with an existing
// personal label will appear in our clients as a shared label. These labels are
// gray by default and will only stay in the shared labels list if there are any
// active tasks with this label.
//
// A user can convert a shared label to a personal label at any time. The label
// will then become customizable and will remain in the account even if not
// assigned to any active tasks.
//
// Shared labels do not appear in the sync response for a user's account. They
// only appear within the labels list of the [Task] that they are assigned to.
//
// You can find more information on the differences between personal and shared
// labels in our [Help Center].
//
// See [Labels] for more details.
//
// [Labels]: https://todoist.com/api/v1/docs#tag/Sync/Labels
// [Help Center]: https://www.todoist.com/help/articles/introduction-to-labels-dSo2eE#shared
type Label struct {
	// The ID of the label.
	ID string `json:"id"`

	// The name of the label.
	Name string `json:"name"`

	// The color of the project icon. Refer to the name column in the
	// https://todoist.com/api/v1/docs#tag/Colors guide for more info.
	Color Color `json:"color"`

	// Label’s order in the label list (a number, where the smallest value should
	// place the label at the top).
	ItemOrder int `json:"item_order"`

	// Whether the label is marked as deleted (a true or false value).
	IsDeleted bool `json:"is_deleted"`

	// Whether the label is a favorite (a true or false value).
	IsFavorite bool `json:"is_favorite"`
}
