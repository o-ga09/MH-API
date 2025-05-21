package di

import (
	"context"
	"mh-api/app/internal/controller"
	healthDriver "mh-api/app/internal/driver/health"
	// "mh-api/app/internal/driver/mysql" // Import will likely be removed
	"mh-api/app/internal/service/health"
)

func InitHealthService(ctx context.Context) *controller.SystemHandler {
	// db := mysql.New(ctx) // Removed
	driver := healthDriver.NewHealthRepository() // DB argument will be removed in a later step
	service := health.NewHealthService(driver)
	h := controller.NewHealthService(*service) // Corrected to assign to new var h

	return &h // Return address of h
}
