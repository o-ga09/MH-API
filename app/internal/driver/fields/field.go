package fields

import (
	"context"
	"fmt"
	"mh-api/app/internal/domain/fields"
	"mh-api/app/internal/driver/mysql"

	"gorm.io/gorm"
)

type fieldRepository struct {
	conn *gorm.DB
}

func NewfieldRepository(conn *gorm.DB) *fieldRepository {
	return &fieldRepository{
		conn: conn,
	}
}

func (r *fieldRepository) Save(ctx context.Context, f fields.Field) error {
	field := mysql.Field{
		FieldId:   f.GetID(),
		MonsterId: f.GetMonsterID(),
		Name:      f.GetName(),
		ImageUrl:  f.GetURL(),
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	err := r.conn.Save(&field).Error
	r.conn.Exec("SET foreign_key_checks = 1")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (r *fieldRepository) Remove(ctx context.Context, fieldId string) error {
	field := mysql.Field{
		FieldId: fieldId,
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	err := r.conn.Delete(&field).Error
	r.conn.Exec("SET foreign_key_checks = 1")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
