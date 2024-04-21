package items

import "context"

type FetchItemByMonster struct {
	qs ItemQueryService
}

func NewFetchItemByMonster(qs ItemQueryService) *FetchItemByMonster {
	return &FetchItemByMonster{qs: qs}
}

func (s *FetchItemByMonster) Run(ctx context.Context, itemId string) (*ItemsByMonster, error) {
	res, error := s.qs.GetItemsByMonster(ctx, itemId)
	if error != nil {
		return nil, error
	}

	return res, nil
}
