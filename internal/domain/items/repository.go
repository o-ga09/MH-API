package items

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	Find(ctx context.Context, params SearchParams) (*SearchResult, error)
	FindAll(ctx context.Context) (Items, error)
	FindByID(ctx context.Context, itemID string) (*Item, error)
	FindByMonsterID(ctx context.Context, monsterID string) (Items, error)
}
