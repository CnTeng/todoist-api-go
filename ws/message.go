package ws

type Message string

type message struct {
	Type Message `json:"type"`
}

const (
	Connected    Message = "connected"
	Disconnected Message = "disconnected"

	Ping            Message = "ping"
	SyncNeeded      Message = "sync_needed"
	CalenderUpdated Message = "calendar_updated"
)
