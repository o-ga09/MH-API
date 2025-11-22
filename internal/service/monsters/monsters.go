package monsters

import (
	"context"
	"mh-api/internal/domain/monsters"
)

//go:generate moq -out monsterservice_mock.go . IMonsterService
type IMonsterService interface {
	FetchMonsterDetail(ctx context.Context, id string) (*FetchMonsterListResult, error)
}

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

func (s *MonsterService) FetchMonsterDetail(ctx context.Context, id string) (*FetchMonsterListResult, error) {
	res, err := s.qs.FetchList(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
