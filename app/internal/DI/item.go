package di

import (
	"context"
	"mh-api/app/internal/controller/item"
	handler "mh-api/app/internal/controller/item"
	"mh-api/app/internal/driver/items"
	"mh-api/app/internal/driver/mysql"
	itemService "mh-api/app/internal/service/item"
)

func InitItemHaandler() *item.ItemHandler {
	db := mysql.New(context.Background())
	repo := items.NewMonsterRepository(db)
	qs := items.NewItemQueryService(db)
	service := itemService.NewItemService(repo, qs)
	handler := handler.NewItemHandler(*service)

	return handler
}
