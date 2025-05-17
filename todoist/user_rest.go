package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

// UserService provides methods for managing user information.
type UserService struct {
	client *Client
}

func NewUserService(client *Client) *UserService {
	return &UserService{client: client}
}

// Get user information.
func (s *UserService) GetUser(ctx context.Context) (*sync.User, error) {
	return get[sync.User](ctx, s.client, UserEndpoint, nil)
}
