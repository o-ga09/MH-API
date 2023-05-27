package store

import (
	"context"
	"fmt"
	"mh-api/api/entity"
	"mh-api/api/interface/repository"

	"gorm.io/gorm"
)

type MonsterRepository struct {
	Db *gorm.DB
}

func NewMonsterRepository(db *gorm.DB) repository.IMonsterRepository {
	return &MonsterRepository{
		Db: db,
	}
}

func (m *MonsterRepository) SelectMonsterAll(ctx context.Context) (entity.Monsters,error) {
	var res entity.Monsters
	result := m.Db.Find(&res)
	if result.RowsAffected == 0 {
		return nil,fmt.Errorf("can not get records: count %d : %w",result.RowsAffected,result.Error)
	}
	return res,nil
}

func (m *MonsterRepository) SelectMonsterById(ctx context.Context, id int) (entity.Monster, error) {
	res := entity.Monster{}
	result := m.Db.Find(&res,entity.Monster{Id: id})
	if result.RowsAffected == 0 {
		return res,fmt.Errorf("can not get record: id = %d : %w",result.RowsAffected,result.Error)
	}
	return res, nil
}