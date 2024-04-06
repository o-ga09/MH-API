package mysql

import (
	"context"
	"log/slog"
	"mh-api/api/config"
	"mh-api/api/middleware"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func New(ctx context.Context) *gorm.DB {
	cfg, err := config.New()
	if err != nil {
		slog.Log(context.Background(), middleware.SeverityError, "environment variable error", "error", err)
	}

	dialector := mysql.Open(cfg.Database_url)

	if db, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		connect(dialector, 100)
	}
	db.Logger = db.Logger.LogMode(logger.Silent)
	slog.Log(ctx, middleware.SeverityInfo, "db connected")
	return db
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if db, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			slog.Log(context.Background(), middleware.SeverityInfo, "db connection retry")
			connect(dialector, count)
			return
		}
		slog.Log(context.Background(), middleware.SeverityInfo, "db connection retry count 100")
		panic(err.Error())
	}
}
