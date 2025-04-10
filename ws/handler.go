package ws

import "context"

type Handler interface {
	HandleNotification(ctx context.Context, noti Notification) error
}
