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

func (qs *itemQueryService) FetchListWithMonster(ctx context.Context) ([]*itemService.FetchItemListWithMonsterDto, error) {
	return nil, nil
}
func (qs *itemQueryService) FetchListByMonster(ctx context.Context) ([]*itemService.FetchItemListByMonsterDto, error) {
	return nil, nil
}
