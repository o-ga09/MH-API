package tribes

import (
	"context"
	Tribes "mh-api/app/internal/domain/tribes"
	"mh-api/app/internal/driver/mysql"
	"strconv"

	"gorm.io/gorm"
)

type tribeRepository struct {
	conn *gorm.DB
}

func NewtribeRepository(conn *gorm.DB) *tribeRepository {
	return &tribeRepository{
		conn: conn,
	}
}

func (r *tribeRepository) Save(ctx context.Context, t Tribes.Tribe) error {
	tribe := mysql.Tribe{
		TribeId:     t.GetID(),
		Name_ja:     t.GetNameJA(),
		Name_en:     t.GetNameEN(),
		Description: t.GetDescription(),
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	err := r.conn.Save(&tribe).Error
	r.conn.Exec("SET foreign_key_checks = 1")
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
	err := r.conn.Delete(&tribe).Error
	if err != nil {
		return err
	}
	return nil
}
