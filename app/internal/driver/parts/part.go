package parts

import (
	"context"
	"mh-api/app/internal/domain/part"
	"mh-api/app/internal/driver/mysql"
)

type partRepository struct{}

func NewMonsterRepository() *partRepository {
	return &partRepository{}
}

func (r *partRepository) Save(ctx context.Context, p part.Part) error {
	data := mysql.Part{
		PartId:      p.GetID(),
		MonsterId:   p.GetMonsterID(),
		Description: p.GetDescription(),
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	err := mysql.CtxFromDB(ctx).Save(&data).Error
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	if err != nil {
		return err
	}
	return nil
}

func (r *partRepository) Remove(ctx context.Context, partId string) error {
	data := mysql.Part{
		PartId: partId,
	}
	err := mysql.CtxFromDB(ctx).Delete(&data).Error
	if err != nil {
		return err
	}
	return nil
}
