package di

import (
	"context"
	"mh-api/app/internal/controller"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/health"
)

func InitHealthService() *controller.SystemHandler {
	db := mysql.New(context.Background())
	driver := mysql.NewHealthRepository(db)
	service := health.NewHealthService(driver)
	handler := controller.NewHealthService(*service)

	return &handler
}
