package di

import (
	"context"
	"mh-api/api/controller/handler"
	"mh-api/api/driver/mysql"
	"mh-api/api/service/monsters"
)

func InitMonstersHandler() *handler.MonsterHandler {
	db := mysql.New(context.Background())
	driver := mysql.NewMonsterRepository(db)
	service := monsters.NewMonsterService(driver)
	handler := handler.NewMonsterHandler(*service)

	return handler
}
