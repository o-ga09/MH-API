package repository

import (
	"context"
	"mh-api/api/usecase/model"
)

type IMonsterRepository interface {
	SelectMonsterAll(ctx context.Context) (model.Monsters, error)
	SelectMonsterById(ctx context.Context, id int) (model.Monster, error)
}