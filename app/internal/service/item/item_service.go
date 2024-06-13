package item

import (
	"context"
	"mh-api/app/internal/domain/items"
)

type ItemService struct {
	repo items.Repository
	qs   ItemQueryService
}

func NewItemService(repo items.Repository, qs ItemQueryService) *ItemService {
	return &ItemService{
		repo: repo,
		qs:   qs,
	}
}

func (s *ItemService) GetItems(ctx context.Context) ([]*FetchItemListDto, error) {
	res, err := s.qs.FetchList(ctx, "")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ItemService) GetItemById(ctx context.Context, id string) ([]*FetchItemListDto, error) {
	res, err := s.qs.FetchList(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ItemService) GetItemsByMonsterList(ctx context.Context, id string) ([]*FetchItemListWithMonsterDto, error) {
	res, err := s.qs.FetchListWithMonster(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ItemService) GetItemByMonsterId(ctx context.Context, id string) ([]*FetchItemListByMonsterDto, error) {
	res, err := s.qs.FetchListByMonster(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
