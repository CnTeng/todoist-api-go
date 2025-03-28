package ws

import "context"

type Handler interface {
	HandleNotification(ctx context.Context, noti Notification) error
}

var DefaultHandler = &defaultHandler{}

type defaultHandler struct{}

func (h *defaultHandler) HandleNotification(ctx context.Context, noti Notification) error {
	return nil
}
