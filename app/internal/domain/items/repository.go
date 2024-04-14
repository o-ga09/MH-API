package items

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	Save(ctx context.Context, m Item) error
	Remove(ctx context.Context, itemId string) error
}
