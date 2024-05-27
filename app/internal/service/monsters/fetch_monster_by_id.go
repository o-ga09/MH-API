package monsters

import (
	"context"
)

type FetchMonsterByIDService struct {
	qs MonsterQueryService
}

func NewFetchMonsterByID(qs MonsterQueryService) *FetchMonsterByIDService {
	return &FetchMonsterByIDService{qs: qs}
}

func (f FetchMonsterByIDService) Run(ctx context.Context, id string) (*FetchMonsterListDto, error) {
	res, err := f.qs.FetchList(ctx, id)
	if err != nil {
		return nil, err
	}

	result := &FetchMonsterListDto{
		Id:                 res[0].Id,
		Name:               res[0].Name,
		Description:        res[0].Description,
		Location:           res[0].Location,
		Category:           res[0].Category,
		Title:              res[0].Title,
		FirstWeak_Attack:   res[0].FirstWeak_Attack,
		SecondWeak_Attack:  res[0].SecondWeak_Attack,
		FirstWeak_Element:  res[0].FirstWeak_Element,
		SecondWeak_Element: res[0].SecondWeak_Element,
		Weakness_attack:    res[0].Weakness_attack,
		Weakness_element:   res[0].Weakness_element,
	}
	return result, nil
}
