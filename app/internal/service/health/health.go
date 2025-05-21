package health

import (
	"context"
	"mh-api/app/internal/domain/health"
)

type HealthService struct {
	repo health.HealthRepository
}

func NewHealthService(repo health.HealthRepository) *HealthService {
	return &HealthService{
		repo: repo,
	}
}

func (s *HealthService) GetStatus(ctx context.Context) error {
	err := s.repo.GetStatus(ctx)
	return err
}
