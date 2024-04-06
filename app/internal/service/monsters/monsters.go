package monsters

import (
	"context"
	"mh-api/app/internal/domain/monsters"
)

type MonsterService struct {
	repo monsters.Repository
}

func NewMonsterService(repo monsters.Repository) *MonsterService {
	return &MonsterService{
		repo: repo,
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
