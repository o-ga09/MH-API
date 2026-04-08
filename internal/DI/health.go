package di

import (
"context"
"mh-api/internal/controller"
"mh-api/internal/database/mysql"
)

func InitHealthService(ctx context.Context) *controller.SystemHandler {
repo := mysql.NewHealthRepository()
h := controller.NewHealthService(repo)
return &h
}
