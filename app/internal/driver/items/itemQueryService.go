package items

import (
	"context"
	param "mh-api/app/internal/controller/item"
	"mh-api/app/internal/driver/mysql"
	itemService "mh-api/app/internal/service/item"
	"strings"

	"gorm.io/gorm"
)

type itemQueryService struct {
	conn *gorm.DB
}

func NewItemQueryService(conn *gorm.DB) *itemQueryService {
	return &itemQueryService{
		conn: conn,
	}
}

func (qs *itemQueryService) FetchList(ctx context.Context, id string) ([]*itemService.FetchItemListDto, error) {
	var item []mysql.Item
	var itemIds []string
	var result *gorm.DB
	var p param.RequestParam
	var err error

	where_clade := ""
	sort := ""

	if id == "" {
		p = ctx.Value(param.ParamKey).(param.RequestParam)
	}

	limit := p.Limit
	offset := p.Offset

	if p.ItemIds != "" {
		itemIds = strings.Split(p.ItemIds, ",")
		where_clade = "item_id IN (?)"
	}

	if p.ItemName != "" && p.ItemIds != "" {
		where_clade += " and name LIKE '%" + p.ItemName + "%' "
	} else if p.ItemName != "" {
		where_clade += " name LIKE '%" + p.ItemName + "%' "
	}

	if p.ItemName == "" && p.ItemNameKana != "" && p.ItemIds != "" {
		where_clade += " and name_kana LIKE '%" + p.ItemNameKana + "%' "
	} else if p.ItemName == "" && p.ItemNameKana != "" {
		where_clade += " name_kana LIKE '%" + p.ItemNameKana + "%' "
	}

	if p.Sort == 1 {
		sort = "item_id DESC"
	} else {
		sort = "item_id ASC"
	}

	if id != "" {
		result = qs.conn.Model(&item).Where("item_id = ? ", id).Find(&item)
	} else if where_clade != "" && p.ItemIds != "" {
		result = qs.conn.Model(&item).Where(where_clade, itemIds).Limit(limit).Offset(offset).Order(sort).Find(&item)
	} else if where_clade != "" {
		result = qs.conn.Model(&item).Where(where_clade).Limit(limit).Offset(offset).Order(sort).Find(&item)
	} else {
		result = qs.conn.Model(&item).Limit(limit).Offset(offset).Order(sort).Find(&item)
	}

	if result.Error != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	res := []*itemService.FetchItemListDto{}
	for _, m := range item {
		r := itemService.FetchItemListDto{
			Id:   m.ItemId,
			Name: m.Name,
			Url:  m.ImageUrl,
		}
		res = append(res, &r)
	}
	return res, err
}

/*
  取得するSQL
  SELECT i.item_id, item.name, item.image_url,
  	( SELECT GROUP_CONCAT(m.monster_id)
	   	FROM item_with_monster AS m
	     WHERE m.item_id = i.item_id
  	) AS monsters
　FROM item_with_monster AS i
  JOIN item AS item ON item.item_id = i.item_id
　GROUP BY i.item_id,item.image_url,item.name;
*/

func (qs *itemQueryService) FetchListWithMonster(ctx context.Context) ([]*itemService.FetchItemListWithMonsterDto, error) {
	var itemIds []string
	var result *gorm.DB
	var p param.RequestParam
	var err error

	r := []struct {
		ItemId   string
		ItemName string
		ImageUrl string
		Monsters string
	}{}

	p = ctx.Value(param.ParamKey).(param.RequestParam)

	where_clade := ""
	sort := ""

	limit := 100
	offset := p.Offset

	if p.ItemIds != "" {
		itemIds = strings.Split(p.ItemIds, ",")
		where_clade = "i.item_id IN (?)"
	}

	if p.ItemName != "" && p.ItemIds != "" {
		where_clade += " and name LIKE '%" + p.ItemName + "%' "
	} else if p.ItemName != "" {
		where_clade += " name LIKE '%" + p.ItemName + "%' "
	}

	if p.ItemName == "" && p.ItemNameKana != "" && p.ItemIds != "" {
		where_clade += " and name_kana LIKE '%" + p.ItemNameKana + "%' "
	} else if p.ItemName == "" && p.ItemNameKana != "" {
		where_clade += " name_kana LIKE '%" + p.ItemNameKana + "%' "
	}

	if p.Sort == 1 {
		sort = "i.item_id DESC"
	} else {
		sort = "i.item_id ASC"
	}

	// サブクエリ
	subQuery := qs.conn.Select("GROUP_CONCAT(m.monster_id)").Where("m.item_id = i.item_id").Table("item_with_monster AS m")

	if where_clade != "" && p.ItemIds != "" {
		result = qs.conn.Table("item_with_monster AS i").Select("i.item_id, item.name AS item_name, item.image_url, (?) AS monsters", subQuery).Joins("JOIN item AS item ON item.item_id = i.item_id").Where(where_clade, itemIds).Group(" i.item_id,image_url,item.name").Limit(limit).Offset(offset).Order(sort).Find(&r)
	} else if where_clade != "" {
		result = qs.conn.Table("item_with_monster AS i").Select("i.item_id, item.name AS item_name, item.image_url, (?) AS monsters", subQuery).Joins("JOIN item AS item ON item.item_id = i.item_id").Where(where_clade).Group(" i.item_id,image_url,item.name").Limit(limit).Offset(offset).Order(sort).Find(&r)
	} else {
		result = qs.conn.Table("item_with_monster AS i").Select("i.item_id, item.name AS item_name, item.image_url, (?) AS monsters", subQuery).Joins("JOIN item AS item ON item.item_id = i.item_id").Group(" i.item_id,image_url,item.name").Limit(limit).Offset(offset).Order(sort).Find(&r)
	}

	if result.Error != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	res := []*itemService.FetchItemListWithMonsterDto{}
	resMonster := []*itemService.Monster{}

	for _, m := range r {
		monsterIds := strings.Split(m.Monsters, ",")
		for _, monster := range monsterIds {
			resMonster = append(resMonster, &itemService.Monster{
				MonsterId: monster,
			})
		}
		r := itemService.FetchItemListWithMonsterDto{
			Id:      m.ItemId,
			Name:    m.ItemName,
			Url:     m.ImageUrl,
			Monster: resMonster,
		}
		res = append(res, &r)
		resMonster = nil
	}
	return res, err
}

/*
	  取得するSQL
	  SELECT i.monster_id, monster_info.name,
	  	( SELECT JSON_ARRAYAGG(m.item_id)
		    FROM item_with_monster AS m
			WHERE m.monster_id = i.monster_id
		) AS monsters
	  FROM item_with_monster AS i
	  JOIN monster AS monster_info ON monster_info.monster_id = i.monster_id
	  GROUP BY i.monster_id, monster_info.name;
*/
func (qs *itemQueryService) FetchListByMonster(ctx context.Context) ([]*itemService.FetchItemListByMonsterDto, error) {
	var item []mysql.ItemWithMonster
	var itemIds []string
	var result *gorm.DB
	var p param.RequestParam
	var err error

	p = ctx.Value(param.ParamKey).(param.RequestParam)

	r := []struct {
		MonsterId   string
		MonsterName string
		Items       []string
	}{}

	where_clade := ""
	sort := ""

	limit := p.Limit
	offset := p.Offset

	if p.ItemIds != "" {
		itemIds = strings.Split(p.ItemIds, ",")
		where_clade = "i.item_id IN (?)"
	}

	if p.ItemName != "" && p.ItemIds != "" {
		where_clade += " and monster_info.name LIKE '%" + p.ItemName + "%' "
	} else if p.ItemName != "" {
		where_clade += " monster_info.name LIKE '%" + p.ItemName + "%' "
	}

	if p.Sort == 1 {
		sort = "i.monster_id DESC"
	} else {
		sort = "i.monster_id ASC"
	}

	// サブクエリ
	subQuery := qs.conn.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Select("JSON_ARRAYAGG(m.item_id)").Where("m.monster_id = i.monster_id").Table("item_with_monster AS m")
	})

	if where_clade != "" && p.ItemIds != "" {
		result = qs.conn.Table("i", &item).Select("i.monster_id, monster_info.name", subQuery).Joins("JOIN monster AS monster_info ON monster_info.monster_id = i.monster_id").Where(where_clade, itemIds).Group(" i.monster_id, monster_info.name").Limit(limit).Offset(offset).Order(sort).Find(&r)
	} else if where_clade != "" {
		result = qs.conn.Table("i", &item).Select("i.monster_id, monster_info.name", subQuery).Joins("JOIN monster AS monster_info ON monster_info.monster_id = i.monster_id").Where(where_clade).Group(" i.monster_id, monster_info.name").Limit(limit).Offset(offset).Order(sort).Find(&r)
	} else {
		result = qs.conn.Table("i", &item).Select("i.monster_id, monster_info.name", subQuery).Joins("JOIN monster AS monster_info ON monster_info.monster_id = i.monster_id").Limit(limit).Group(" i.monster_id, monster_info.name").Offset(offset).Order(sort).Find(&r)
	}

	if result.Error != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	res := []*itemService.FetchItemListByMonsterDto{}
	resItem := []*itemService.FetchItemListDto{}

	for _, m := range r {
		for _, item := range m.Items {
			resItem = append(resItem, &itemService.FetchItemListDto{
				Id: item,
			})
		}
		r := itemService.FetchItemListByMonsterDto{
			MonsterId:   m.MonsterId,
			MonsterName: m.MonsterName,
			Item:        resItem,
		}
		res = append(res, &r)
	}
	return res, err
}
