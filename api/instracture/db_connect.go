package instracture

import (
	"context"
	"database/sql"
	"fmt"
	"mh-api/api/config"
	"time"

	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	db, err := sql.Open("mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?parseTime=true",
				cfg.DBUser, cfg.DBPassword,
				cfg.DBHost, cfg.DBPort,
				cfg.DBName,
			),
		)

	if err != nil {
		return nil, nil, err
	}

	ctx, cansel := context.WithTimeout(ctx,2*time.Second)
	defer cansel()

	if err := db.PingContext(ctx); err != nil {
		return nil, func() {_ = db.Close()}, err
	}

	xdb := sqlx.NewDb(db,"mysql")
	return xdb, func() {_ = db.Close()}, nil

}