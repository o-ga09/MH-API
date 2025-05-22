package weakness

import "context"

type Repository interface {
	Save(ctx context.Context, m Weakness) error
	Remove(ctx context.Context, monsterId string) error
}
