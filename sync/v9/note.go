package sync

// ItemNote represents a note in the Todoist application.
type ItemNote struct {
	// ID is the ID of the note.
	ID string `json:"id"`

	// PostedUID is the ID of the user that posted the note.
	PostedUID string `json:"posted_uid"`

	// ItemID is the item which the note is part of.
	ItemID string `json:"item_id"`

	// Content is the content of the note. This value may contain markdown-formatted text and hyperlinks.
	Content string `json:"content"`

	// FileAttachment is a file attached to the note.
	FileAttachment *FileAttachment `json:"file_attachment,omitempty"`

	// UIDsToNotify is a list of user IDs to notify.
	UIDsToNotify []string `json:"uids_to_notify"`

	// IsDeleted indicates whether the note is marked as deleted.
	IsDeleted bool `json:"is_deleted"`

	// PostedAt is the date when the note was posted.
	PostedAt string `json:"posted_at"`

	// Reactions is a list of emoji reactions and corresponding user IDs.
	Reactions *Reactions `json:"reactions,omitempty"`
}

// ItemNote represents a note in the Todoist application.
type ProjectNote struct {
	// ID is the ID of the note.
	ID string `json:"id"`

	// PostedUID is the ID of the user that posted the note.
	PostedUID string `json:"posted_uid"`

	// ItemID is the item which the note is part of.
	ProjectID string `json:"project_id"`

	// Content is the content of the note. This value may contain markdown-formatted text and hyperlinks.
	Content string `json:"content"`

	// FileAttachment is a file attached to the note.
	FileAttachment *FileAttachment `json:"file_attachment,omitempty"`

	// UIDsToNotify is a list of user IDs to notify.
	UIDsToNotify []string `json:"uids_to_notify"`

	// IsDeleted indicates whether the note is marked as deleted.
	IsDeleted bool `json:"is_deleted"`

	// PostedAt is the date when the note was posted.
	PostedAt string `json:"posted_at"`

	// Reactions is a list of emoji reactions and corresponding user IDs.
	Reactions *Reactions `json:"reactions,omitempty"`
}

type FileAttachment struct {
	FileName    string `json:"file_name"`
	FileSize    int    `json:"file_size"`
	FileType    string `json:"file_type"`
	FileURL     string `json:"file_url"`
	Image       string `json:"image"`
	ImageHeight int    `json:"image_height"`
	ImageWidth  int    `json:"image_width"`
	UploadState string `json:"upload_state"`
}

type Reactions map[string][]string
