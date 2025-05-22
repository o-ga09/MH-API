package health

import "context"

type HealthRepository interface {
	GetStatus(ctx context.Context) error
}
