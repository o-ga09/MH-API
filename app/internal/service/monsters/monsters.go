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

func (s *MonsterService) FetchMonsterDetail(ctx context.Context, id string) ([]*FetchMonsterListDto, error) {
	res, err := s.qs.FetchList(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *MonsterService) SaveMonsters(ctx context.Context, param MonsterDto) error {
	saveData := monsters.NewMonster(param.ID, param.Name, param.Description, param.Element)
	err := s.repo.Save(ctx, saveData)
	if err != nil {
		return err
	}

	return nil
}

func (s *MonsterService) RemoveMonsters(ctx context.Context, id string) error {
	err := s.repo.Remove(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
