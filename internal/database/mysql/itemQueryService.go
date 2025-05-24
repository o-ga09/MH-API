package mysql

import (
	"context"
	"errors"
	"fmt"
	"mh-api/internal/domain/items"

	"gorm.io/gorm"
)

type itemQueryService struct{}

func NewItemQueryService() items.Repository {
	return &itemQueryService{}
}

func (s *itemQueryService) FindAll(ctx context.Context) (items.Items, error) {
	db := CtxFromDB(ctx)

	var itemModels []Item
	if err := db.Find(&itemModels).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch items: %w", err)
	}

	var domainItems items.Items
	for _, model := range itemModels {
		domainItem := items.NewItem(model.ItemId, model.Name, model.ImageUrl)
		domainItems = append(domainItems, *domainItem)
	}

	return domainItems, nil
}

func (s *itemQueryService) FindByID(ctx context.Context, itemID string) (*items.Item, error) {
	db := CtxFromDB(ctx)

	var itemModel Item
	if err := db.Where("item_id = ?", itemID).First(&itemModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to fetch item by ID: %w", err)
	}

	return items.NewItem(itemModel.ItemId, itemModel.Name, itemModel.ImageUrl), nil
}

func (s *itemQueryService) FindByMonsterID(ctx context.Context, monsterID string) (items.Items, error) {
	db := CtxFromDB(ctx)

	var itemModels []Item
	if err := db.Where("monster_id = ?", monsterID).Find(&itemModels).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch items by monster ID: %w", err)
	}

	var domainItems items.Items
	for _, model := range itemModels {
		domainItem := items.NewItem(model.ItemId, model.Name, model.ImageUrl)
		domainItems = append(domainItems, *domainItem)
	}

	return domainItems, nil
}
