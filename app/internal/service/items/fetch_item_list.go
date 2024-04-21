package items

import "context"

type FetchItemList struct {
	qs ItemQueryService
}

func NewFetchItemList(qs ItemQueryService) *FetchItemList {
	return &FetchItemList{qs: qs}
}

func (s *FetchItemList) Run(ctx context.Context) ([]*ItemDto, error) {
	res, error := s.qs.GetItems(ctx)
	if error != nil {
		return nil, error
	}

	return res, nil
}
