package mysql

import (
	"context"
	"fmt"
	"mh-api/internal/domain/items"
)

type ItemModel struct {
	ItemID   string `gorm:"column:item_id;primaryKey"`
	ItemName string `gorm:"column:item_name"`
}

type itemQueryService struct{}

func NewItemQueryService() items.Repository {
	return &itemQueryService{}
}

func (s *itemQueryService) FindAll(ctx context.Context) (items.Items, error) {
	db := CtxFromDB(ctx)
	if db == nil {
		return nil, fmt.Errorf("database connection not found in context")
	}

	var itemModels []ItemModel
	if err := db.Find(&itemModels).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch items: %w", err)
	}

	var domainItems items.Items
	for _, model := range itemModels {
		domainItem := items.NewItem(model.ItemID, model.ItemName, "")
		domainItems = append(domainItems, *domainItem)
	}

	return domainItems, nil
}

func (s *itemQueryService) Save(ctx context.Context, m items.Item) error {
	return fmt.Errorf("Save method not implemented in itemQueryService")
}

func (s *itemQueryService) Remove(ctx context.Context, itemId string) error {
	return fmt.Errorf("Remove method not implemented in itemQueryService")
}
