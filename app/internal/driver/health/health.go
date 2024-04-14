package mysql

import "gorm.io/gorm"

type healthRepository struct {
	conn *gorm.DB
}

func NewHealthRepository(conn *gorm.DB) *healthRepository {
	return &healthRepository{conn: conn}
}

func (h *healthRepository) GetStatus() error {
	db, err := h.conn.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}
