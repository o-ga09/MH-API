package music

import (
	"context"
	"mh-api/app/internal/domain/music"
	"mh-api/app/internal/driver/mysql"

	"gorm.io/gorm"
)

type musicRepository struct {
	conn *gorm.DB
}

func NewmusicRepository(conn *gorm.DB) *musicRepository {
	return &musicRepository{
		conn: conn,
	}
}

func (r *musicRepository) Get(ctx context.Context, monsterId string) (music.Musics, error) {
	m := []mysql.Music{}
	err := r.conn.Find(&m).Error
	if err != nil {
		return nil, err
	}

	res := music.Musics{}
	for _, r := range m {
		res = append(res, *music.NewMusic(r.MusicId, r.MonsterId, r.Name, r.ImageUrl))
	}

	return res, nil
}

func (r *musicRepository) Save(ctx context.Context, m music.Music) error {
	data := mysql.Music{
		MusicId:   m.GetID(),
		MonsterId: m.GetMonsterID(),
		Name:      m.GetName(),
		ImageUrl:  m.GetURL(),
	}
	r.conn.Exec("SET foreign_key_checks = 0")
	err := r.conn.Save(&data).Error
	r.conn.Exec("SET foreign_key_checks = 1")
	if err != nil {
		return err
	}
	return nil
}

func (r *musicRepository) Remove(ctx context.Context, musicId string) error {
	data := mysql.Music{
		MusicId: musicId,
	}
	err := r.conn.Delete(&data).Error
	if err != nil {
		return err
	}
	return nil
}
