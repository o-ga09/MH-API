package di

import (
	"context"
	"mh-api/app/internal/controller/monster"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/driver/queryservice"
	"mh-api/app/internal/driver/repository"
	"mh-api/app/internal/service/monsters"
)

func ProvideMonsterHandler(ctx context.Context) *monster.MonsterHandler {
	db := mysql.New(ctx)
	repo := repository.NewMonsterRepository(db)
	qs := queryservice.NewMonsterQueryService(db)
	fetchListService := monsters.NewFetchMonsterListService(qs)
	fetchByIdService := monsters.NewFetchMonsterByID(qs)
	fetchRankingService := monsters.NewFetchMonsterRankingService(qs)
	saveMonsterService := monsters.NewSaveMonsterService(repo)
	removeMonsterService := monsters.NewRemoveMonsterService(repo)
	monsterHandler := monster.NewMonsterHandler(*fetchListService, *fetchByIdService, *fetchRankingService, *saveMonsterService, *removeMonsterService)
	return monsterHandler
}
