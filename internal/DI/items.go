package di

import (
	"context"
	itemController "mh-api/internal/controller/item"
	"mh-api/internal/database/mysql"
)

func InitItemsHandler(ctx context.Context) *itemController.ItemHandler {
	itemRepo := mysql.NewItemQueryService()
	monsterRepo := mysql.NewMonsterRepository()
	return itemController.NewItemHandler(itemRepo, monsterRepo)
}
