package di

import (
	"context"
	"mh-api/app/internal/controller"
	healthDriver "mh-api/app/internal/driver/health"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/health"
)

func InitHealthService(ctx context.Context) *controller.SystemHandler {
	db := mysql.New(ctx)
	driver := healthDriver.NewHealthRepository(db)
	service := health.NewHealthService(driver)
	handler := controller.NewHealthService(*service)

	return &handler
}
