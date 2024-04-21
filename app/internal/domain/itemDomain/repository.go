package itemdomain

import "context"

//go:generate moq -out repository_mock.go . ItemRepository
type ItemRepository interface {
	Save(ctx context.Context, m Item) error
	Remove(ctx context.Context, itemId string) error
}
