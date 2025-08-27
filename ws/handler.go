package ws

import "context"

// Handler defines the interface for handling WebSocket messages.
type Handler interface {
	// HandleMessage handles incoming WebSocket messages.
	HandleMessage(ctx context.Context, msg Message) error
}
