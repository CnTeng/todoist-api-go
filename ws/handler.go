package ws

import "context"

type Handler interface {
	HandleMessage(ctx context.Context, msg Message) error
}
