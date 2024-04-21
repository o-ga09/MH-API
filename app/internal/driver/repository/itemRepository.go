package repository

import (
	"context"
	itemdomain "mh-api/app/internal/domain/itemDomain"
	"mh-api/app/internal/driver/schema"

	"gorm.io/gorm"
)

type ItemRepository struct {
	conn *gorm.DB
}

func NewItemRepository(conn *gorm.DB) *ItemRepository {
	return &ItemRepository{conn}
}

func (r *ItemRepository) Save(ctx context.Context, m itemdomain.Item) error {
	item := schema.Item{
		ItemId:      m.ItemID(),
		MonsterId:   m.MonsterID(),
		Name:        m.ItemName(),
		Description: m.ItemDescription(),
		ImageUrl:    m.ItemImageURL(),
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	if err := r.conn.Create(&item).Error; err != nil {
		return err
	}
	r.conn.Exec("SET foreign_key_checks = 1")

	return nil
}

func (r *ItemRepository) Remove(ctx context.Context, itemId string) error {
	r.conn.Exec("SET foreign_key_checks = 0")
	if err := r.conn.Delete(&schema.Item{ItemId: itemId}).Error; err != nil {
		return err
	}
	r.conn.Exec("SET foreign_key_checks = 1")

	return nil
}
