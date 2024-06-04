package di

import (
	"context"
	"mh-api/app/internal/controller/item"
	handler "mh-api/app/internal/controller/item"
	"mh-api/app/internal/driver/mysql"
)

func InitItemHaandler() *item.ItemHandler {
	db := mysql.New(context.Background())

	handler := handler.NewItemHandler(service)

	return handler
}
