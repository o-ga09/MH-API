package di

import (
	"context"
	handler "mh-api/app/internal/controller/monster"
	monsterDriver "mh-api/app/internal/driver/monsters"

	"mh-api/app/internal/service/monsters"
)

func InitMonstersHandler(ctx context.Context) *handler.MonsterHandler {
	// db := mysql.New(ctx) // Removed
	repo := monsterDriver.NewMonsterRepository()
	qs := monsterDriver.NewmonsterQueryService()
	service := monsters.NewMonsterService(repo, qs)
	h := handler.NewMonsterHandler(*service)

	return h
}
