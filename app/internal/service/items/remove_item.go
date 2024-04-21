package items

import (
	"context"
	itemdomain "mh-api/app/internal/domain/itemDomain"
)

type RemoveItem struct {
	repo itemdomain.ItemRepository
}

func NewRemoveItem(repo itemdomain.ItemRepository) *RemoveItem {
	return &RemoveItem{repo: repo}
}

func (s *RemoveItem) Run(ctx context.Context, itemId string) error {
	err := s.repo.Remove(ctx, itemId)
	if err != nil {
		return err
	}

	return nil
}
