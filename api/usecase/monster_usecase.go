package usecase

import (
	"mh-api/api/domain/servise"
	"mh-api/api/usecase/model"

	"golang.org/x/net/context"
)

type MonsterUsecase interface {
	FindAllMonsters(ctx context.Context) (model.Monsters, error)
	FindMonsterById(ctx context.Context,id int) (model.Monster,error)
}

type monsterUsecase struct {
	use servise.MonsterService
}

func NewMonsterUsecase(u servise.MonsterService) MonsterUsecase {
	return &monsterUsecase{
		use: u,
	}
}

func (u *monsterUsecase) FindAllMonsters(ctx context.Context) (model.Monsters,error) {
	return u.use.FindAllMonsters(ctx)
}

func (u *monsterUsecase) FindMonsterById(ctx context.Context, id int) (model.Monster,error) {
	return u.use.FindMonsterById(ctx,id)
}