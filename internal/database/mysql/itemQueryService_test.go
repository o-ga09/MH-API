package mysql

import (
	"context"
	"mh-api/internal/domain/items" // ドメイン層のitemsパッケージをインポート
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemQueryService_Save(t *testing.T) {
	ctx := context.Background()
	service := NewItemQueryService() // items.Repository を返す

	// ダミーのItemオブジェクトを作成
	dummyItem := items.NewItem("dummyId", "dummyName", "dummyUrl")

	err := service.Save(ctx, *dummyItem)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Save method not implemented")
}

func TestItemQueryService_Remove(t *testing.T) {
	ctx := context.Background()
	service := NewItemQueryService() // items.Repository を返す

	err := service.Remove(ctx, "dummyId")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Remove method not implemented")
}

// TestItemQueryService_FindAll は、itemsテーブルのマイグレーションと
// テストデータ準備が整うまでスキップします。
// func TestItemQueryService_FindAll(t *testing.T) {
// 	// TODO: DBマイグレーションとテストデータ準備後に実装
// }
