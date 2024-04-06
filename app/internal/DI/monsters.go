package di

import (
	"context"
	handler "mh-api/app/internal/controller/monster"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/monsters"
)

func InitMonstersHandler() *handler.MonsterHandler {
	db := mysql.New(context.Background())
	driver := mysql.NewMonsterRepository(db)
	service := monsters.NewMonsterService(driver)
	handler := handler.NewMonsterHandler(*service)

	return handler
}
