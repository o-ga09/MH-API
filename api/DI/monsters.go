package di

import (
	"context"
	"mh-api/api/controller/handler"
	"mh-api/api/driver"
	"mh-api/api/gateway"
	"mh-api/api/service"
)

func InitMonstersHandler() *handler.MonsterHandler {
	db := driver.New(context.Background())
	driver := driver.ProvideMonsterDriver(db)
	gateway := gateway.ProvideMonsterDriver(driver)
	service := service.ProvideMonsterDriver(gateway)
	handler := handler.ProvideMonsterHandler(service)

	return &handler
}