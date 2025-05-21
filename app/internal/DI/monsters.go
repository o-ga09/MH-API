package di

import (
	"context"
	handler "mh-api/app/internal/controller/monster"
	monsterDriver "mh-api/app/internal/driver/monsters"
	// "mh-api/app/internal/driver/mysql" // mysql import no longer needed here
	"mh-api/app/internal/service/monsters"
)

func InitMonstersHandler(ctx context.Context) *handler.MonsterHandler {
	// db := mysql.New(ctx) // Removed
	repo := monsterDriver.NewMonsterRepository() // DB argument will be removed in a later step
	qs := monsterDriver.NewmonsterQueryService()   // DB argument will be removed in a later step
	service := monsters.NewMonsterService(repo, qs)
	h := handler.NewMonsterHandler(*service) // Corrected variable name from handler to h

	return h
}
