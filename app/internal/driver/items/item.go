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

func (r *itemRepository) Get(ctx context.Context, itemId string) (items.Items, error) {
	var items items.Items
	if err := r.conn.Find(&items).Error; err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return items, nil
}

func (r *itemRepository) Save(ctx context.Context, m items.Item) error {
	i := mysql.Item{
		ItemId:   m.GetID(),
		Name:     m.GetName(),
		ImageUrl: m.GetURL(),
	}
	if err := r.conn.Save(&i).Error; err != nil {
		return fmt.Errorf(err.Error())
	}
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
