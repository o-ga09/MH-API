package di

import (
	"context"
	"mh-api/internal/controller/weapon"
	"mh-api/internal/database/mysql"
)

func InitWeaponHandler(ctx context.Context) *weapon.WeaponHandler {
	repo := mysql.NewWeaponRepository()
	return weapon.NewWeaponHandler(repo)
}
