package di

import (
	"context"
	"mh-api/app/internal/controller"
	healthDriver "mh-api/app/internal/driver/health"

	"mh-api/app/internal/service/health"
)

func InitHealthService(ctx context.Context) *controller.SystemHandler {
	driver := healthDriver.NewHealthRepository()
	service := health.NewHealthService(driver)
	h := controller.NewHealthService(*service)

	return &h
}
