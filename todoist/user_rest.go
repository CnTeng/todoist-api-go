package todoist

import (
	"context"

	"github.com/CnTeng/todoist-api-go/sync"
)

func (c *Client) GetUser(ctx context.Context) (*sync.User, error) {
	return get[sync.User](ctx, c, UserEndpoint, nil)
}
