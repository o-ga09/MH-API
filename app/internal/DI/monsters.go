package di

import (
	"context"
	handler "mh-api/app/internal/controller/monster"
	monsterDriver "mh-api/app/internal/driver/monsters"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/monsters"
)

func InitMonstersHandler() *handler.MonsterHandler {
	db := mysql.New(context.Background())
	repo := monsterDriver.NewMonsterRepository(db)
	qs := monsterDriver.NewmonsterQueryService(db)
	service := monsters.NewMonsterService(repo, qs)
	handler := handler.NewMonsterHandler(*service)

	return handler
}
