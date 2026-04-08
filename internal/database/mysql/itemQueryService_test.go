package mysql

import (
	"context"
	"errors"
	"testing"

	"mh-api/internal/domain/items"

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

	testItems := createItemData(t, ctx)

	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{name: "正常系: アイテムを全件取得できる", want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &itemQueryService{}
			got, err := service.FindAll(ctx)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Len(t, got, tt.want)
			for i, item := range got {
				assert.Equal(t, testItems[i].ItemId, item.ItemId)
				assert.Equal(t, testItems[i].Name, item.Name)
				assert.Equal(t, testItems[i].ImageUrl, item.ImageUrl)
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

	testItems := createItemData(t, ctx)

	tests := []struct {
		name    string
		itemID  string
		want    *items.Item
		wantErr bool
		errType error
	}{
		{
			name:   "正常系: 存在するIDの場合",
			itemID: testItems[0].ItemId,
			want:   testItems[0],
		},
		{
			name:    "異常系: 存在しないIDの場合",
			itemID:  "non-existent-id",
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
		{
			name:    "異常系: 空のIDの場合",
			itemID:  "",
			wantErr: true,
			errType: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &itemQueryService{}
			got, err := service.FindByID(ctx, tt.itemID)

			if tt.wantErr {
				require.Error(t, err)
				if tt.errType != nil {
					assert.True(t, errors.Is(err, tt.errType))
				}
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, tt.want.ItemId, got.ItemId)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.ImageUrl, got.ImageUrl)
		})
	}
}

func createItemData(t *testing.T, ctx context.Context) []*items.Item {
	t.Helper()

	itemList := []*items.Item{
		{ItemId: "0000000001", Name: "ポーション", ImageUrl: "images/potion.png"},
		{ItemId: "0000000002", Name: "グレートポーション", ImageUrl: "images/great_potion.png"},
		{ItemId: "0000000003", Name: "メガポーション", ImageUrl: "images/mega_potion.png"},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(itemList).Error)

	return itemList
}
