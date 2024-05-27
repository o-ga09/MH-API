package monsterdomain

import "context"

//go:generate moq -out repository_mock.go . MonsterRepository
type MonsterRepository interface {
	Save(ctx context.Context, m *Monster) error
	Remove(ctx context.Context, monsterId string) error
}
