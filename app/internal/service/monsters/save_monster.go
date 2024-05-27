package monsters

import (
	"context"
	monsterdomain "mh-api/app/internal/domain/monsterDomain"
)

type SaveMonsterService struct {
	repo monsterdomain.MonsterRepository
}

func NewSaveMonsterService(repo monsterdomain.MonsterRepository) *SaveMonsterService {
	return &SaveMonsterService{repo: repo}
}

func (s SaveMonsterService) Run(ctx context.Context, m *MonsterDto) error {
	monster := monsterdomain.NewMonster(m.Id, m.Name, m.Description)
	err := s.repo.Save(ctx, monster)
	if err != nil {
		return err
	}
	return nil
}
