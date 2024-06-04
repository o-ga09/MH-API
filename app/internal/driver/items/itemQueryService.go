package items

import (
	"context"
	itemService "mh-api/app/internal/service/item"

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
	return nil, nil
}
func (qs *itemQueryService) FetchListWithMonster(ctx context.Context) ([]*itemService.FetchItemListWithMonsterDto, error) {
	return nil, nil
}
func (qs *itemQueryService) FetchListByMonster(ctx context.Context) ([]*itemService.FetchItemListByMonsterDto, error) {
	return nil, nil
}
