package sync

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	Lang         string `json:"lang"`
	WebsocketURL string `json:"websocket_url"`
	IsDeleted    bool   `json:"is_deleted"`
}
