package sync

import "context"

const getCompletedEndpoint = baseURL + "/completed/get_all"

func (c *Client) GetCompletedInfo(ctx context.Context, params *CompletedGetParams) (*CompletedGetResponse, error) {
	return do[CompletedGetParams, CompletedGetResponse](ctx, c, getCompletedEndpoint, params)
}
