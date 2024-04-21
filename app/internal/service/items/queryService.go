package items

import "context"

//go:generate moq -out repository_mock.go . ItemQueryService
type ItemQueryService interface {
	GetItems(ctx context.Context) ([]*ItemDto, error)
	GetItem(ctx context.Context, itemId string) (*ItemDto, error)
	GetItemsByMonster(ctx context.Context, itemId string) (*ItemsByMonster, error)
}
