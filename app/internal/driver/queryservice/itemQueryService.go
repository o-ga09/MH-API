package queryservice

import (
	"context"
	"mh-api/app/internal/driver/schema"
	"mh-api/app/internal/service/items"

	"gorm.io/gorm"
)

type ItemQueryService struct {
	conn *gorm.DB
}

func NewItemQueryService(conn *gorm.DB) *ItemQueryService {
	return &ItemQueryService{conn}
}

func (s *ItemQueryService) GetItems(ctx context.Context) ([]*items.ItemDto, error) {
	var item []schema.Item

	result := s.conn.Find(&item)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var res []*items.ItemDto
	for _, i := range item {
		res = append(res, &items.ItemDto{
			ID:          i.ItemId,
			MonsterId:   i.MonsterId,
			Name:        i.Name,
			Description: i.Description,
			ImageURL:    i.ImageUrl,
		})
	}

	return res, nil
}

func (s *ItemQueryService) GetItem(ctx context.Context, itemId string) (*items.ItemDto, error) {
	var item schema.Item

	result := s.conn.Where(&schema.Item{ItemId: itemId}).First(&item)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &items.ItemDto{
		ID:          item.ItemId,
		MonsterId:   item.MonsterId,
		Name:        item.Name,
		Description: item.Description,
		ImageURL:    item.ImageUrl,
	}, nil
}

func (s *ItemQueryService) GetItemsByMonster(ctx context.Context, itemId string) (*items.ItemsByMonster, error) {
	itemsByMonster := []struct {
		MonsterID   string
		MonsterName string
		ItemID      string
		ItemName    string
	}{}

	result := s.conn.Table("item").Select("monster.monster_id, monster.name as monster_name, item.item_id, item.name as item_name").Joins("JOIN monster on item.monster_id = monster.monster_id").Where("item_id = ?", itemId).Scan(&itemsByMonster)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var res items.ItemsByMonster
	res.ItemId = itemsByMonster[0].ItemID
	res.ItemName = itemsByMonster[0].ItemName
	monsters := []items.Monster{}
	for _, i := range itemsByMonster {
		monsters = append(monsters, items.Monster{
			ID:   i.MonsterID,
			Name: i.MonsterName,
		})
	}
	res.Monster = monsters
	return &res, nil
}
