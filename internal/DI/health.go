package di

import (
	"context"
	"mh-api/internal/controller"
	"mh-api/internal/database/mysql"

	"mh-api/internal/service/health"
)

func InitHealthService(ctx context.Context) *controller.SystemHandler {
	driver := mysql.NewHealthRepository()
	service := health.NewHealthService(driver)
	h := controller.NewHealthService(*service)

	return &h
}
