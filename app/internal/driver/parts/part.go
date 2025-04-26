package parts

import (
	"context"
	"mh-api/app/internal/domain/part"
	"mh-api/app/internal/driver/mysql"

	"gorm.io/gorm"
)

type partRepository struct {
	conn *gorm.DB
}

func NewMonsterRepository(conn *gorm.DB) *partRepository {
	return &partRepository{
		conn: conn,
	}
}

func (r *partRepository) Save(ctx context.Context, p part.Part) error {
	data := mysql.Part{
		PartId:      p.GetID(),
		MonsterId:   p.GetMonsterID(),
		Description: p.GetDescription(),
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	err := r.conn.Save(&data).Error
	r.conn.Exec("SET foreign_key_checks = 1")
	if err != nil {
		return err
	}
	return nil
}

func (r *partRepository) Remove(ctx context.Context, partId string) error {
	data := mysql.Part{
		PartId: partId,
	}
	err := r.conn.Delete(&data).Error
	if err != nil {
		return err
	}
	return nil
}
