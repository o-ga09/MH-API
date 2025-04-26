package items

import (
	"context"
	"fmt"
	"mh-api/app/internal/domain/items"
	"mh-api/app/internal/driver/mysql"

	"gorm.io/gorm"
)

type itemRepository struct {
	conn *gorm.DB
}

func NewMonsterRepository(conn *gorm.DB) *itemRepository {
	return &itemRepository{
		conn: conn,
	}
}

func (r *itemRepository) Save(ctx context.Context, m items.Item) error {
	i := mysql.Item{
		ItemId:   m.GetID(),
		Name:     m.GetName(),
		ImageUrl: m.GetURL(),
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	if err := r.conn.Save(&i).Error; err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	r.conn.Exec("SET foreign_key_checks = 1")
	return nil
}

func (r *itemRepository) Remove(ctx context.Context, itemId string) error {
	item := mysql.Item{
		ItemId: itemId,
	}
	err := r.conn.Delete(&item).Error
	if err != nil {
		return err
	}
	return nil
}
