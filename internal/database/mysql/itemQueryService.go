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
	var itemList []*items.Item
	if err := CtxFromDB(ctx).Find(&itemList).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch items: %w", err)
	}
	return itemList, nil
}

func (s *itemQueryService) FindByID(ctx context.Context, itemID string) (*items.Item, error) {
	var item items.Item
	if err := CtxFromDB(ctx).Where("item_id = ?", itemID).First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to fetch item by ID: %w", err)
	}
	return &item, nil
}

func (s *itemQueryService) FindByMonsterID(ctx context.Context, monsterID string) (items.Items, error) {
	var itemList []*items.Item
	if err := CtxFromDB(ctx).Where("monster_id = ?", monsterID).Find(&itemList).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch items by monster ID: %w", err)
	}
	return itemList, nil
}
