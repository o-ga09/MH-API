package monsters

import (
	"context"
	monsterdomain "mh-api/app/internal/domain/monsterDomain"
)

type RemoveMonsterService struct {
	repo monsterdomain.MonsterRepository
}

func NewRemoveMonsterService(repo monsterdomain.MonsterRepository) *RemoveMonsterService {
	return &RemoveMonsterService{repo: repo}
}

func (r RemoveMonsterService) Run(ctx context.Context, monsterId string) error {
	err := r.repo.Remove(ctx, monsterId)
	if err != nil {
		return err
	}
	return nil
}
