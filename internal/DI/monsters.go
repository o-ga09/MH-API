package di

import (
	"context"
	handler "mh-api/internal/controller/monster"
	"mh-api/internal/database/mysql"

	"mh-api/internal/service/monsters"
)

func InitMonstersHandler(ctx context.Context) *handler.MonsterHandler {
	// db := mysql.New(ctx) // Removed
	repo := mysql.NewMonsterRepository()
	qs := mysql.NewmonsterQueryService()
	service := monsters.NewMonsterService(repo, qs)
	h := handler.NewMonsterHandler(*service)

	return h
}
