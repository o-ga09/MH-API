package monsters

import "context"

//go:generate moq -out queryService_mock.go . MonsterQueryService
type MonsterQueryService interface {
	FetchList(ctx context.Context, id string) ([]*FetchMonsterListDto, error)
	FetchRank(ctx context.Context) ([]*FetchMonsterRankingDto, error)
}
