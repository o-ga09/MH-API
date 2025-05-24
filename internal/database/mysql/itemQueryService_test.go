package mysql

import (
	"context"
	"errors"
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
	tx := db.Begin()
	defer tx.Rollback()

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

func TestItemQueryService_FindByID(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	tx := db.Begin()
	defer tx.Rollback()

	// テストデータを準備
	testItems := createItemData(t, ctx)

	// テストケースを定義
	tests := []struct {
		name    string
		itemID  string
		want    *items.Item
		wantErr bool
		errType error
	}{
		{
			name:    "正常系: 存在するIDの場合",
			itemID:  testItems[0].GetID(),
			want:    &testItems[0],
			wantErr: false,
			errType: nil,
		},
		{
			name:    "異常系: 存在しないIDの場合",
			itemID:  "non-existent-id",
			want:    nil,
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
		{
			name:    "異常系: 空のIDの場合",
			itemID:  "",
			want:    nil,
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// クエリサービスの初期化
			service := &itemQueryService{}

			// テスト対象メソッド実行
			got, err := service.FindByID(ctx, tt.itemID)

			// アサーション
			if tt.wantErr {
				require.Error(t, err)
				if tt.errType != nil {
					assert.True(t, errors.Is(err, tt.errType), "expected error type: %v, got: %v", tt.errType, err)
				}
			} else {
				require.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tt.want.GetID(), got.GetID())
				assert.Equal(t, tt.want.GetName(), got.GetName())
				assert.Equal(t, tt.want.GetURL(), got.GetURL())
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
