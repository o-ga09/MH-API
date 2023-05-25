package store

import (
	"context"
	"fmt"
	"mh-api/api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(ctx context.Context, cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBName, cfg.DBPassword,
		cfg.DBHost, cfg.DBPort,
		cfg.DBName,
	)),&gorm.Config{})
	if err != nil {
		return nil,fmt.Errorf("err: can not open DB connect: %w",err)
	}

	return db,  nil
}