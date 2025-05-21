package fields

import (
	"context"
	"fmt"
	"mh-api/app/internal/domain/fields"
	"mh-api/app/internal/driver/mysql"
)

type fieldRepository struct{}

func NewfieldRepository() *fieldRepository {
	return &fieldRepository{}
}

func (r *fieldRepository) Save(ctx context.Context, f fields.Field) error {
	field := mysql.Field{
		FieldId:   f.GetID(),
		MonsterId: f.GetMonsterID(),
		Name:      f.GetName(),
		ImageUrl:  f.GetURL(),
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	err := mysql.CtxFromDB(ctx).Save(&field).Error
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (r *fieldRepository) Remove(ctx context.Context, fieldId string) error {
	field := mysql.Field{
		FieldId: fieldId,
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	err := mysql.CtxFromDB(ctx).Delete(&field).Error
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
