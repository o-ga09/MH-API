package service

import (
	"mh-api/api/entity"
	"mh-api/api/interface/monster"

	"golang.org/x/net/context"
)

type MonsterService interface {
	FindAllMonsters(ctx context.Context) (entity.Monsters, error)
	FindMonsterById(ctx context.Context,id int) (entity.Monster,error)
}

type monsterService struct {
	use monster.IMonsterService
}

func NewMonsterUsecase(u monster.IMonsterService) MonsterService {
	return &monsterService{
		use: u,
	}
}

func (u *monsterService) FindAllMonsters(ctx context.Context) (entity.Monsters,error) {
	return u.use.FindAllMonsters(ctx)
}

func (u *monsterService) FindMonsterById(ctx context.Context, id int) (entity.Monster,error) {
	return u.use.FindMonsterById(ctx,id)
}