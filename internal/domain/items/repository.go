package items

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	FindAll(ctx context.Context) (Items, error)
}
