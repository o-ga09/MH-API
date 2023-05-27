package store

import (
	"context"
	"fmt"
	"mh-api/api/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func New(ctx context.Context, cfg *config.Config) (*gorm.DB, error) {
	dialector := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword,
		cfg.DBHost, cfg.DBName,
	))
	var err error
	if db,err = gorm.Open(dialector,&gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		connect(dialector,100)
	}
	return db,  nil
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if db, err = gorm.Open(dialector,&gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())
	}
}