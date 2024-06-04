package music

import (
	"context"
	"fmt"
	param "mh-api/app/internal/controller/music"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/music"
	"strings"

	"gorm.io/gorm"
)

type musicQueryService struct {
	conn *gorm.DB
}

func NewmusicQueryService(conn *gorm.DB) *musicQueryService {
	return &musicQueryService{
		conn: conn,
	}
}

func (q *musicQueryService) FetchList(ctx context.Context, id string) ([]*music.FetchMusicListDto, error) {
	var bgm []mysql.Music
	var bgmIds []string
	var result *gorm.DB
	var p param.RequestParam
	var err error

	where_clade := ""
	sort := ""

	if id == "" {
		p = ctx.Value("param").(param.RequestParam)
	}

	limit := p.Limit
	offset := p.Offset

	if p.BgmIds != "" {
		bgmIds = strings.Split(p.BgmIds, ",")
		where_clade = "music_id IN (?)"
	}

	if p.BgmName != "" && p.BgmIds != "" {
		where_clade += " and name LIKE '%" + p.BgmName + "%' "
	} else if p.BgmName != "" {
		where_clade += " name LIKE '%" + p.BgmName + "%' "
	}

	if p.Sort == "1" {
		sort = "music_id ASC"
	} else {
		sort = "music_id DESC"
	}

	fmt.Println("id: ", id)
	if id != "" {
		result = q.conn.Debug().Model(&bgm).Preload("BgmRanking").Where("music_id = ? ", id).Find(&bgm)
	} else if where_clade != "" && p.BgmIds != "" {
		result = q.conn.Model(&bgm).Preload("BgmRanking").Where(where_clade, bgmIds).Limit(limit).Offset(offset).Order(sort).Find(&bgm)
	} else if where_clade != "" {
		result = q.conn.Model(&bgm).Preload("BgmRanking").Where(where_clade).Limit(limit).Offset(offset).Order(sort).Find(&bgm)
	} else {
		result = q.conn.Model(&bgm).Preload("BgmRanking").Limit(limit).Offset(offset).Order(sort).Find(&bgm)
	}

	if result.Error != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	res := []*music.FetchMusicListDto{}
	for _, m := range bgm {
		r := music.FetchMusicListDto{
			Id:   m.MusicId,
			Name: m.Name,
			Url:  m.Url,
		}
		res = append(res, &r)
	}
	return res, err
}
func (q *musicQueryService) FetchRank(ctx context.Context) ([]*music.FetchMusicRankingDto, error) {
	var bgm []mysql.Music
	var musicIds []string
	var result *gorm.DB
	var p param.RequestParam
	var err error

	r := []struct {
		MusicId  string `gorm:"column:music_id"`
		Name     string `gorm:"column:name"`
		Url      string `gorm:"column:url"`
		Ranking  string `gorm:"column:ranking"`
		VoteYear string `gorm:"column:vote_year"`
	}{}

	where_clade := ""
	sort := ""

	p = ctx.Value("param").(param.RequestParam)

	limit := p.Limit
	offset := p.Offset

	if p.BgmIds != "" {
		musicIds = strings.Split(p.BgmIds, ",")
		where_clade = "music.music_id IN (?)"
	}

	if p.BgmName != "" && p.BgmIds != "" {
		where_clade += " and music.name LIKE '%" + p.BgmName + "%' "
	} else if p.BgmName != "" {
		where_clade += " music.name LIKE '%" + p.BgmName + "%' "
	}

	if p.Sort == "1" {
		sort = "music.music_id ASC"
	} else {
		sort = "music.music_id DESC"
	}

	if where_clade != "" && p.BgmIds != "" {
		result = q.conn.Debug().Model(&bgm).Select("music.music_id as music_id", "music.name as name", "music.url as url", "bgm_ranking.ranking as ranking", "bgm_ranking.vote_year as vote_year").Joins("JOIN bgm_ranking ON bgm_ranking.music_id = music.music_id").Where(where_clade, musicIds).Limit(limit).Offset(offset).Order(sort).Find(&r)
	} else if where_clade != "" {
		result = q.conn.Model(&bgm).Select("music.music_id as music_id", "music.name as name", "music.url as url", "bgm_ranking.ranking as ranking", "bgm_ranking.vote_year as vote_year").Joins("JOIN bgm_ranking ON bgm_ranking.music_id = music.music_id").Where(where_clade).Limit(limit).Offset(offset).Order(sort).Find(&r)
	} else {
		result = q.conn.Model(&bgm).Select("music.music_id as music_id", "music.name as name", "music.url as url", "bgm_ranking.ranking as ranking", "bgm_ranking.vote_year as vote_year").Joins("JOIN bgm_ranking ON bgm_ranking.music_id = music.music_id").Limit(limit).Offset(offset).Order(sort).Find(&r)
	}

	if result.Error != nil {
		return nil, err
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	res := []*music.FetchMusicRankingDto{}
	grouped := make(map[string][]*music.Ranking)

	for _, m := range r {
		grouped[m.MusicId] = append(grouped[m.MusicId], &music.Ranking{
			Rank:     m.Ranking,
			VoteYear: m.VoteYear,
		})
	}
	for _, m := range r {

		r := music.FetchMusicRankingDto{
			Id:      m.MusicId,
			Name:    m.Name,
			Url:     m.Url,
			Ranking: grouped[m.MusicId],
		}
		res = append(res, &r)
	}
	return res, err
}
