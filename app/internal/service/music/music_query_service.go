package music

import "context"

//go:generate moq -out queryService_mock.go . MonsterQueryService
type MusicQueryService interface {
	FetchList(ctx context.Context, id string) ([]*FetchMusicListDto, error)
	FetchRank(ctx context.Context) ([]*FetchMusicRankingDto, error)
}

type FetchMusicListDto struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type FetchMusicRankingDto struct {
	Id      string     `json:"id,omitempty"`
	Name    string     `json:"name,omitempty"`
	Url     string     `json:"url,omitempty"`
	Ranking []*Ranking `json:"ranking,omitempty"`
}

type Ranking struct {
	Rank     string `json:"rank,omitempty"`
	VoteYear string `json:"vote_year,omitempty"`
}
