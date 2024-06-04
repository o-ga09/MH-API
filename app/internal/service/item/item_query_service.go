package item

import "context"

//go:generate moq -out queryService_mock.go . MonsterQueryService
type ItemQueryService interface {
	FetchList(ctx context.Context, id string) ([]*FetchItemListDto, error)
	FetchListWithMonster(ctx context.Context) ([]*FetchItemListWithMonsterDto, error)
	FetchListByMonster(ctx context.Context) ([]*FetchItemListByMonsterDto, error)
}

type FetchItemListDto struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type FetchItemListWithMonsterDto struct {
	Id      string     `json:"id,omitempty"`
	Name    string     `json:"name,omitempty"`
	Url     string     `json:"url,omitempty"`
	Ranking []*Ranking `json:"ranking,omitempty"`
}

type FetchItemListByMonsterDto struct {
	Id      string     `json:"id,omitempty"`
	Name    string     `json:"name,omitempty"`
	Url     string     `json:"url,omitempty"`
	Ranking []*Ranking `json:"ranking,omitempty"`
}

type Ranking struct {
	Rank     string `json:"rank,omitempty"`
	VoteYear string `json:"vote_year,omitempty"`
}
