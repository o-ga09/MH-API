package di

import (
	"context"
	"mh-api/internal/controller/weapon"
	"mh-api/internal/database/mysql"
	weapons_service "mh-api/internal/service/weapons"
)

// InitWeaponHandler は WeaponHandler とその依存関係を初期化し、返します。
func InitWeaponHandler(ctx context.Context) *weapon.WeaponHandler {
	weaponQueryService := mysql.NewWeaponQueryService()
	weaponService := weapons_service.NewWeaponService(weaponQueryService)
	weaponHandler := weapon.NewWeaponHandler(weaponService)

	return weaponHandler
}
