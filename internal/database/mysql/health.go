package mysql

import (
	"context"
	"errors"
)

type healthRepository struct{}

func NewHealthRepository() *healthRepository {
	return &healthRepository{}
}

func (h *healthRepository) GetStatus(ctx context.Context) error {
	db := CtxFromDB(ctx)
	if db == nil {
		return errors.New("database connection not found")
	}
	conn, err := db.DB()
	if err != nil {
		return err
	}
	return conn.Ping()
}
