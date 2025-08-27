package sync

// User represents a Todoist user.
type User struct {
	// The user's ID.
	ID string `json:"id"`

	// The user's email.
	Email string `json:"email"`

	// The user's real name formatted as Firstname Lastname.
	FullName string `json:"full_name"`

	// The user's language, which can take one of the following values:
	//   da, de, en, es, fi, fr, it, ja, ko, nl, pl, pt_BR, ru, sv, tr, zh_CN,
	//   zh_TW.
	Lang string `json:"lang"`

	// The websocket URL for the user.
	WebsocketURL string `json:"websocket_url"`

	// Whether the user is marked as deleted (a true or false value).
	IsDeleted bool `json:"is_deleted"`
}
