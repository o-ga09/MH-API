package fields

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	Save(ctx context.Context, m Field) error
	Remove(ctx context.Context, monsterId string) error
}
