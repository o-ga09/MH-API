package monsters

import (
	"context"
	"mh-api/app/internal/domain/monsters"
)

type MonsterService struct {
	repo monsters.Repository
	qs   MonsterQueryService
}

func NewMonsterService(repo monsters.Repository, qs MonsterQueryService) *MonsterService {
	return &MonsterService{
		repo: repo,
		qs:   qs,
	}
}

func (s *MonsterService) GetMonster(ctx context.Context, id string) ([]*MonsterDto, error) {
	m, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	var res []*MonsterDto
	for _, r := range m {
		res = append(res, &MonsterDto{
			ID:          r.GetId(),
			Description: r.GetDesc(),
			Name:        r.GetName(),
		})
	}

	return res, nil
}

func (s *MonsterService) FetchMonsterDetail(ctx context.Context, id string) ([]*FetchMonsterListDto, error) {
	res, err := s.qs.FetchMonsterList(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *MonsterService) FetchMonsterRanking(ctx context.Context) ([]*FetchMonsterRankingDto, error) {
	res, err := s.qs.FetchMonsterRanking(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
