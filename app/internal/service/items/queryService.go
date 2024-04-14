package items

import "context"

type ItemQueryService interface {
	GetItems(ctx context.Context) ([]*ItemDto, error)
	GetItem(ctx context.Context, itemId string) (*ItemDto, error)
	GetItemsByMonster(ctx context.Context) (*ItemsByMonster, error)
}

type ItemService struct {
	qs ItemQueryService
}

func NewItemQueryService(qs ItemQueryService) *ItemService {
	return &ItemService{qs: qs}
}

func (s *ItemService) GetItems(ctx context.Context) ([]*ItemDto, error) {
	res, error := s.qs.GetItems(ctx)
	if error != nil {
		return nil, error
	}

	return res, nil
}

func (s *ItemService) GetItem(ctx context.Context, itemId string) (*ItemDto, error) {
	res, error := s.qs.GetItem(ctx, itemId)
	if error != nil {
		return nil, error
	}

	return res, nil
}

func (s *ItemService) GetItemsByMonster(ctx context.Context) (*ItemsByMonster, error) {
	res, error := s.qs.GetItemsByMonster(ctx)
	if error != nil {
		return nil, error
	}

	return res, nil
}
