package armors

import "context"

//go:generate moq -out repository_mock.go . Repository
type Repository interface {
	GetAll(ctx context.Context) (Armors, error)
	GetByID(ctx context.Context, armorId string) (*Armor, error)
}
