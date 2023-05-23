package servise

import (
	"context"
	"mh-api/api/domain/repository"
	"mh-api/api/usecase/model"
)

type MonsterService interface {
	FindAllMonsters(ctx context.Context) (model.Monsters,error)
	FindMonsterById(ctx context.Context,id int) (model.Monster, error)
}

type monsterService struct {
	repo repository.IMonsterRepository
}

func NewMosterService(r repository.IMonsterRepository) MonsterService {
	return &monsterService{
		repo: r,
	}
}

func (s *monsterService) FindAllMonsters(ctx context.Context) (model.Monsters, error) {
	return s.repo.SelectMonsterAll(ctx)
}

func (s *monsterService) FindMonsterById(ctx context.Context, id int) (model.Monster, error) {
	return s.repo.SelectMonsterById(ctx,id)
}