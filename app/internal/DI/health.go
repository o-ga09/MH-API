package di

import (
	"context"
	"mh-api/app/internal/controller"
	healthDriver "mh-api/app/internal/driver/health"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/health"
)

func InitHealthService() *controller.SystemHandler {
	db := mysql.New(context.Background())
	driver := healthDriver.NewHealthRepository(db)
	service := health.NewHealthService(driver)
	handler := controller.NewHealthService(*service)

	return &handler
}
