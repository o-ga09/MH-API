package items

import (
	"context"
	"fmt"
	"mh-api/app/internal/domain/items"
	"mh-api/app/internal/driver/mysql"
)

type itemRepository struct{}

func NewMonsterRepository() *itemRepository {
	return &itemRepository{}
}

func (r *itemRepository) Save(ctx context.Context, m items.Item) error {
	i := mysql.Item{
		ItemId:   m.GetID(),
		Name:     m.GetName(),
		ImageUrl: m.GetURL(),
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 0")
	if err := mysql.CtxFromDB(ctx).Save(&i).Error; err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	mysql.CtxFromDB(ctx).Exec("SET foreign_key_checks = 1")
	return nil
}

func (r *itemRepository) Remove(ctx context.Context, itemId string) error {
	item := mysql.Item{
		ItemId: itemId,
	}
	err := mysql.CtxFromDB(ctx).Delete(&item).Error
	if err != nil {
		return err
	}
	return nil
}
