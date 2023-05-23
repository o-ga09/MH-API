package mysql

import (
	"context"
	"database/sql"
	"mh-api/api/domain/repository"
	"mh-api/api/usecase/model"
)

type MonsterRepository struct {
	Db *sql.DB
}

func NewMonsterRepository(db *sql.DB) repository.IMonsterRepository {
	return &MonsterRepository{
		Db: db,
	}
}

func (m *MonsterRepository) SelectMonsterAll(ctx context.Context) (model.Monsters,error) {
	//TBD:
	var TBD model.Monsters
	return TBD,nil
}

func (m *MonsterRepository) SelectMonsterById(ctx context.Context, id int) (model.Monster, error) {
	//TBD:
	var TBD model.Monster
	return TBD,nil
}