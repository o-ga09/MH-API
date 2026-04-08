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
	gormDB := CtxFromDB(ctx)
	if gormDB == nil {
		return errors.New("database connection not found")
	}
	conn, err := gormDB.DB()
	if err != nil {
		return err
	}
	return conn.Ping()
}
