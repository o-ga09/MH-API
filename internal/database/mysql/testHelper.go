package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"mh-api/pkg/config"
	"os"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var testDB *gorm.DB

// setupTestDB はテスト用のDBセットアップを行うヘルパー関数
func setupTestDB(ctx context.Context) context.Context {
	os.Setenv("DATABASE_URL", "mh-api:P@ssw0rd@tcp(127.0.0.1:3306)/ci?charset=utf8&parseTime=True&loc=Local")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// テスト用DBの接続情報
	dsn := cfg.Database_url

	// テスト用DBに接続
	testDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.WithValue(ctx, CtxKey, testDB)

	// シードデータ投入
	migrationTestData()
	return ctx
}

// テスト用のDBのマイグレーション
func migrationTestData() {
	// マイグレーション
	migrations := &migrate.FileMigrationSource{
		Dir: "../../../db/migrations",
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied %d migrations\n", n)

	// TRUNCATE文のSQLの読み込み
	truncateSQL, err := os.ReadFile("../../../db/seed/00_truncate.sql")
	if err != nil {
		log.Fatal(err)
	}
	// テーブルのクリーンアップ
	statements := strings.SplitSeq(string(truncateSQL), ";")
	// テーブルのクリーンアップ
	for stmt := range statements {
		// 空の文を除外
		if strings.TrimSpace(stmt) != "" {
			if err := testDB.Exec(stmt).Error; err != nil {
				log.Fatal(err)
			}
		}
	}
}

// CtxFromDB はコンテキストからDBを取得するヘルパー関数
func CtxFromTestDB(ctx context.Context) *gorm.DB {
	return testDB
}
