package di

import (
	"context"
	handler "mh-api/app/internal/controller/music"
	musicDriver "mh-api/app/internal/driver/music"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/music"
)

func InitBGMHandler() *handler.BGMHandler {
	db := mysql.New(context.Background())
	repo := musicDriver.NewMusicRepository(db)
	qs := musicDriver.NewmusicQueryService(db)
	service := music.NewMusicService(repo, qs)
	handler := handler.NewBGMHandler(*service)

	return handler
}
