package health

import "mh-api/app/internal/domain/health"

type HealthService struct {
	repo health.HealthRepository
}

func NewHealthService(repo health.HealthRepository) *HealthService {
	return &HealthService{
		repo: repo,
	}
}

func (s *HealthService) GetStatus() error {
	err := s.repo.GetStatus()
	return err
}
