package monsters

import "context"

type FetchMonsterListService struct {
	qs MonsterQueryService
}

func NewFetchMonsterListService(qs MonsterQueryService) *FetchMonsterListService {
	return &FetchMonsterListService{qs: qs}
}

func (s *FetchMonsterListService) Run(ctx context.Context, id string) ([]*FetchMonsterListDto, error) {
	res, err := s.qs.FetchList(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]*FetchMonsterListDto, len(res))
	for i, r := range res {
		result[i] = &FetchMonsterListDto{
			Id:                 r.Id,
			Name:               r.Name,
			Description:        r.Description,
			Location:           r.Location,
			Category:           r.Category,
			Title:              r.Title,
			FirstWeak_Attack:   r.FirstWeak_Attack,
			SecondWeak_Attack:  r.SecondWeak_Attack,
			FirstWeak_Element:  r.FirstWeak_Element,
			SecondWeak_Element: r.SecondWeak_Element,
			Weakness_attack:    r.Weakness_attack,
			Weakness_element:   r.Weakness_element,
		}
	}
	return result, nil
}
