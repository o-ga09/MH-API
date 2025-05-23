package mysql

import (
	"context"
	"fmt"
	"mh-api/internal/domain/items" // ドメイン層のitemsパッケージをインポート

	"gorm.io/gorm" // GORMをインポート
)

// ItemModel はデータベースのitemsテーブルのスキーマを表す構造体 (仮定)
// TODO: 実際のテーブルスキーマに合わせて調整が必要
type ItemModel struct {
	ItemID   string `gorm:"column:item_id;primaryKey"` // item_id カラム
	ItemName string `gorm:"column:item_name"`          // item_name カラム
	// 他に必要なカラムがあればここに追加
}

// TableName はItemModelが対応するテーブル名を返すメソッド
func (ItemModel) TableName() string {
	return "items" // 仮のテーブル名
}

// itemQueryService は items.Repository インターフェースを実装する構造体
type itemQueryService struct{}

// NewItemQueryService は itemQueryService の新しいインスタンスを生成します
func NewItemQueryService() items.Repository {
	return &itemQueryService{}
}

// FindAll は全てのアイテム情報を取得します
func (s *itemQueryService) FindAll(ctx context.Context) (items.Items, error) {
	db := CtxFromDB(ctx)
	if db == nil {
		return nil, fmt.Errorf("database connection not found in context")
	}

	var itemModels []ItemModel
	// TODO: エラーハンドリングをより丁寧に行う (例: gorm.ErrRecordNotFound の扱い)
	if err := db.Find(&itemModels).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch items: %w", err)
	}

	var domainItems items.Items
	for _, model := range itemModels {
		// ItemModel からドメインの Item 構造体に変換
		// 現状の items.Item は imageUrl も引数に取るため、空文字を設定
		domainItem := items.NewItem(model.ItemID, model.ItemName, "") 
		domainItems = append(domainItems, *domainItem)
	}

	return domainItems, nil
}

// Save はアイテム情報を保存します (未実装)
func (s *itemQueryService) Save(ctx context.Context, m items.Item) error {
	return fmt.Errorf("Save method not implemented in itemQueryService")
}

// Remove はアイテム情報を削除します (未実装)
func (s *itemQueryService) Remove(ctx context.Context, itemId string) error {
	return fmt.Errorf("Remove method not implemented in itemQueryService")
}
