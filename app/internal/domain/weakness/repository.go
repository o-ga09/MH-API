package weakness

import "context"

type Repository interface {
	get(ctx context.Context, monsterId string) (Weakness error)
	save(ctx context.Context, m Weakness) error
	remove(ctx context.Context, monsterId string) error
}
