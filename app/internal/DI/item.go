package di

import (
	"context"
	"mh-api/app/internal/controller/item"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/driver/queryservice"
	"mh-api/app/internal/driver/repository"
	itemService "mh-api/app/internal/service/items"
)

func ProvideItemHandler(ctx context.Context) *item.ItemHandler {
	db := mysql.New(ctx)
	repo := repository.NewItemRepository(db)
	qs := queryservice.NewItemQueryService(db)
	fetchItemList := itemService.NewFetchItemList(qs)
	fetchItemByMonster := itemService.NewFetchItemByMonster(qs)
	fetchItemById := itemService.NewFetchItemById(qs)
	saveItem := itemService.NewSaveItem(repo)
	removeItem := itemService.NewRemoveItem(repo)
	return item.NewItemHandler(*fetchItemList, *fetchItemByMonster, *fetchItemById, *saveItem, *removeItem)
}
