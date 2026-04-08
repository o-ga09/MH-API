package di

import (
"context"
handler "mh-api/internal/controller/monster"
"mh-api/internal/database/mysql"
)

func InitMonstersHandler(ctx context.Context) *handler.MonsterHandler {
repo := mysql.NewMonsterRepository()
return handler.NewMonsterHandler(repo)
}
