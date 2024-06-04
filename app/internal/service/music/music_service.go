package music

import (
	"context"
	"mh-api/app/internal/domain/music"
)

type MusicService struct {
	repo music.Repository
	qs   MusicQueryService
}

func NewMusicService(repo music.Repository, qs MusicQueryService) *MusicService {
	return &MusicService{
		repo: repo,
		qs:   qs,
	}
}

func (s *MusicService) FetchList(ctx context.Context, id string) ([]*FetchMusicListDto, error) {
	res, err := s.qs.FetchList(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MusicService) FetchRank(ctx context.Context) ([]*FetchMusicRankingDto, error) {
	res, err := s.qs.FetchRank(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
