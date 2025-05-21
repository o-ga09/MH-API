package tribes

import (
	"context"
	Tribes "mh-api/app/internal/domain/tribes"
	"mh-api/app/internal/driver/mysql"
	"strconv"

	"gorm.io/gorm"
)

type tribeRepository struct {
}

func NewtribeRepository() *tribeRepository {
	return &tribeRepository{}
}

func (r *tribeRepository) Save(ctx context.Context, t Tribes.Tribe) error {
	tribe := mysql.Tribe{
		TribeId:     t.GetID(),
		Name_ja:     t.GetNameJA(),
		Name_en:     t.GetNameEN(),
		Description: t.GetDescription(),
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	err := mysql.CtxFromDB(ctx).Save(&tribe).Error
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	if err != nil {
		return err
	}
	return nil
}

func (r *tribeRepository) Remove(ctx context.Context, Id string) error {
	i, _ := strconv.Atoi(Id)
	tribe := mysql.Tribe{
		Model: gorm.Model{ID: uint(i)},
	}
	err := mysql.CtxFromDB(ctx).Delete(&tribe).Error
	if err != nil {
		return err
	}
	return nil
}
