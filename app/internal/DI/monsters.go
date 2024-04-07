package di

import (
	"context"
	handler "mh-api/app/internal/controller/monster"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/monsters"
)

func InitMonstersHandler() *handler.MonsterHandler {
	db := mysql.New(context.Background())
	repo := mysql.NewMonsterRepository(db)
	qs := mysql.NewmonsterQueryService(db)
	service := monsters.NewMonsterService(repo, qs)
	handler := handler.NewMonsterHandler(*service)

	return handler
}
