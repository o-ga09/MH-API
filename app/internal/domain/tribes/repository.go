package Tribes

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	Save(ctx context.Context, m Tribe) error
	Remove(ctx context.Context, monsterId string) error
}