package monsters

import "context"

type FetchMonsterRankingService struct {
	qs MonsterQueryService
}

func NewFetchMonsterRankingService(qs MonsterQueryService) *FetchMonsterRankingService {
	return &FetchMonsterRankingService{qs: qs}
}

func (f FetchMonsterRankingService) Run(ctx context.Context) ([]*FetchMonsterRankingDto, error) {
	res, err := f.qs.FetchRank(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*FetchMonsterRankingDto, len(res))
	for i, r := range res {
		result[i] = &FetchMonsterRankingDto{
			Id:          r.Id,
			Name:        r.Name,
			Description: r.Description,
			Location:    r.Location,
			Category:    r.Category,
			Title:       r.Title,
			Ranking:     r.Ranking,
		}
	}
	return result, nil
}
