package di

import (
	"context"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/monsters"
)

func BatchDI() *monsters.MonsterService {
	db := mysql.New(context.Background())
	repo := mysql.NewMonsterRepository(db)
	qs := mysql.NewmonsterQueryService(db)
	return monsters.NewMonsterService(repo, qs)
}
