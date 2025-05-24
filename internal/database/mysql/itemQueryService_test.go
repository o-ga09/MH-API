package mysql

import (
	"context"
	"mh-api/internal/domain/items"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestItemQueryService_FindAll(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	db.Begin()
	defer db.Rollback()

	// テストデータを準備
	testItems := createItemData(t, ctx)

	// テストケースを定義
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{
			name:    "正常系: アイテムを全件取得できる",
			want:    3,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &itemQueryService{}

			got, err := service.FindAll(ctx)

			assert.True(t, (err != nil) == tt.wantErr)
			if !tt.wantErr {
				assert.Len(t, got, tt.want)
				for i, item := range got {
					assert.Equal(t, testItems[i].GetID(), item.GetID())
					assert.Equal(t, testItems[i].GetName(), item.GetName())
					assert.Equal(t, testItems[i].GetURL(), item.GetURL())
				}
			}
		})
	}
}

func createItemData(t *testing.T, ctx context.Context) items.Items {
	t.Helper()

	itemModels := []Item{
		{ItemId: "0000000001", Name: "ポーション", ImageUrl: "images/potion.png"},
		{ItemId: "0000000002", Name: "グレートポーション", ImageUrl: "images/great_potion.png"},
		{ItemId: "0000000003", Name: "メガポーション", ImageUrl: "images/mega_potion.png"},
	}

	err := CtxFromTestDB(ctx).Create(itemModels).Error
	require.NoError(t, err)

	var domainItems items.Items
	for _, model := range itemModels {
		domainItem := items.NewItem(model.ItemId, model.Name, model.ImageUrl)
		domainItems = append(domainItems, *domainItem)
	}

	return domainItems
}
