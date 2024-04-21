package items

import (
	"context"
	itemdomain "mh-api/app/internal/domain/itemDomain"
)

type SaveItem struct {
	repo itemdomain.ItemRepository
}

func NewSaveItem(repo itemdomain.ItemRepository) *SaveItem {
	return &SaveItem{repo: repo}
}

func (s *SaveItem) Run(ctx context.Context, item *ItemDto) error {
	i := itemdomain.NewItem(item.ID, item.MonsterId, item.Name, item.Description, item.ImageURL)
	err := s.repo.Save(ctx, *i)
	if err != nil {
		return err
	}

	return nil
}
