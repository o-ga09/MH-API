package monster

import (
	"context"
	"mh-api/api/entity"
	"mh-api/api/interface/repository"
)

type IMonsterService interface {
	FindAllMonsters(ctx context.Context) (entity.Monsters,error)
	FindMonsterById(ctx context.Context,id int) (entity.Monster, error)
}

type monster struct {
	repo repository.IMonsterRepository
}

func NewMosterService(r repository.IMonsterRepository) IMonsterService {
	return &monster{
		repo: r,
	}
}

func (s *monster) FindAllMonsters(ctx context.Context) (entity.Monsters, error) {
	return s.repo.SelectMonsterAll(ctx)
}

func (s *monster) FindMonsterById(ctx context.Context, id int) (entity.Monster, error) {
	return s.repo.SelectMonsterById(ctx,id)
}