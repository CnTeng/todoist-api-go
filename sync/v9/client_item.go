package sync

import (
	"github.com/CnTeng/todoist-api-go/internal/utils"
	"github.com/google/go-querystring/query"
)

const (
	quickAddEndpoint    = baseURL + "/quick/add"
	getItemInfoEndpoint = baseURL + "/items/get"
)

func (sc *SyncClient) AddItem(args *ItemAddArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) UpdateItem(args *ItemUpdateArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) MoveItem(args *ItemMoveArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) ReorderItems(args *ItemReorderArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) DeleteItem(args *ItemDeleteArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) CompleteItem(args *ItemCompleteArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) UncompleteItem(args *ItemUncompleteArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) CompleteRecurringItem(args *ItemCompleteRecurringArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) CloseItem(args *ItemCloseArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) UpdateItemDayOrders(args *ItemUpdateDayOrdersArgs) (*Response, error) {
	cs := &Commands{NewCommand(args)}
	return sc.Write(cs)
}

func (sc *SyncClient) GetItemInfo(params *ItemGetInfoParams) (*ItemGetInfoResponse, error) {
	p, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[ItemGetInfoResponse](sc.client, getItemInfoEndpoint, sc.token)
	return r.WithParameters(p).Post()
}

func (sc *SyncClient) QuickAddItem(params *ItemQuickAddParams) (*Item, error) {
	p, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	r := utils.NewRequest[Item](sc.client, quickAddEndpoint, sc.token)
	return r.WithParameters(p).Post()
}
