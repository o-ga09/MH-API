package mysql

import (
	"context"
	"fmt"
	"mh-api/internal/domain/items"
)

type itemQueryService struct{}

func NewItemQueryService() items.Repository {
	return &itemQueryService{}
}

func (s *itemQueryService) FindAll(ctx context.Context) (items.Items, error) {
	db := CtxFromDB(ctx)
	if db == nil {
		return nil, fmt.Errorf("database connection not found in context")
	}

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
