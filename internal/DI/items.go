package di

import (
	"context"

	itemController "mh-api/internal/controller/item"
	"mh-api/internal/database/mysql"
	itemService "mh-api/internal/service/items"
)

// InitItemsHandler は ItemHandler とその依存関係を初期化し、ItemHandler のインスタンスを返します。
func InitItemsHandler(ctx context.Context) *itemController.ItemHandler {
	// 1. リポジトリ層の初期化
	itemRepo := mysql.NewItemQueryService() // NewItemQueryService は items.Repository を返す想定

	// 2. サービス層の初期化
	itemsSvc := itemService.NewService(itemRepo)

	// 3. コントローラー層（ハンドラー）の初期化
	itemCtrl := itemController.NewItemHandler(itemsSvc)

	return itemCtrl
}
