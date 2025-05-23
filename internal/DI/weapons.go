package di

import (
	"context"
	"mh-api/internal/controller/weapon"
	"mh-api/internal/database/mysql"
	weapons_service "mh-api/internal/service/weapons"
)

// InitWeaponHandler は WeaponHandler とその依存関係を初期化し、返します。
func InitWeaponHandler(ctx context.Context) *weapon.WeaponHandler {
	// 1. mysql.New を呼び出して、DBインスタンスが設定されたコンテキストを取得
	//    New は内部で once.Do を使用しているため、DBの初期化は一度だけ行われる。
	//    ここで渡す ctx は、アプリケーションのベースコンテキスト (例: context.Background()) で良い。
	//    server.go から渡される ctx をそのまま使う。
	dbInitializedCtx := mysql.New(ctx)

	// 2. DBインスタンスが設定されたコンテキストから *gorm.DB を取得
	dbInstance := mysql.CtxFromDB(dbInitializedCtx)

	// 3. 取得した DB インスタンスを使用して依存関係を構築
	weaponQueryService := mysql.NewWeaponQueryService(dbInstance)
	weaponService := weapons_service.NewWeaponService(weaponQueryService)
	weaponHandler := weapon.NewWeaponHandler(weaponService)

	return weaponHandler
}
