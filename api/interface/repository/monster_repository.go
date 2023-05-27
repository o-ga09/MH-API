package repository

import (
	"context"
	"mh-api/api/entity"
)

type IMonsterRepository interface {
	SelectMonsterAll(ctx context.Context) (entity.Monsters, error)
	SelectMonsterById(ctx context.Context, id int) (entity.Monster, error)
}