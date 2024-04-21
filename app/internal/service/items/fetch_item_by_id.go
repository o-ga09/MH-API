package items

import "context"

type FetchItemById struct {
	qs ItemQueryService
}

func NewFetchItemById(qs ItemQueryService) *FetchItemById {
	return &FetchItemById{qs: qs}
}

func (s *FetchItemById) Run(ctx context.Context, id string) (*ItemDto, error) {
	result, err := s.qs.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
