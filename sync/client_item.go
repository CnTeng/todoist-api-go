package sync

import "context"

const (
	quickAddEndpoint = baseURL + "/quick/add"
	getItemsEndpoint = baseURL + "/items/get"
)

func (c *Client) AddItem(ctx context.Context, args *ItemAddArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateItem(ctx context.Context, args *ItemUpdateArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) MoveItem(ctx context.Context, args *ItemMoveArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) ReorderItems(ctx context.Context, args *ItemReorderArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteItem(ctx context.Context, args *ItemDeleteArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) DeleteItems(ctx context.Context, args []*ItemDeleteArgs) (*SyncResponse, error) {
	cmds := make(Commands, 0, len(args))
	for _, arg := range args {
		cmds = append(cmds, NewCommand(arg))
	}
	return c.executeCommands(ctx, &cmds)
}

func (c *Client) CompleteItem(ctx context.Context, args *ItemCompleteArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UncompleteItem(ctx context.Context, args *ItemUncompleteArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) CompleteRecurringItem(ctx context.Context, args *ItemCompleteRecurringArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) CloseItem(ctx context.Context, args *ItemCloseArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) UpdateItemDayOrders(ctx context.Context, args *ItemUpdateDayOrdersArgs) (*SyncResponse, error) {
	return c.executeCommand(ctx, args)
}

func (c *Client) GetItemInfo(ctx context.Context, params *ItemGetInfoParams) (*ItemGetInfoResponse, error) {
	return do[ItemGetInfoParams, ItemGetInfoResponse](ctx, c, getItemsEndpoint, params)
}

func (c *Client) QuickAddItem(ctx context.Context, params *ItemQuickAddParams) (*Item, error) {
	return do[ItemQuickAddParams, Item](ctx, c, quickAddEndpoint, params)
}
