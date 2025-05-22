package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"strings"

	"mh-api/pkg/config"
	"mh-api/pkg/constant"
	"os"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var testDB *gorm.DB

func BeforeTest() {
	var err error
	os.Setenv("DATABASE_URL", "root:pass@tcp(127.0.0.1:3306)/ci?charset=utf8&parseTime=True&loc=Local")
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		slog.Log(context.Background(), constant.SeverityError, "environment variable error", "error", err)
	}
	dialector := mysql.Open(cfg.Database_url)

	var db *gorm.DB

	if db, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		connect(ctx, dialector)
	}

	weak1 := []*Weakness{
		{MonsterId: "0000000001", PartId: "0001", Fire: "45", Water: "45", Lightning: "45", Ice: "45", Dragon: "45", Slashing: "45", Blow: "45", Bullet: "45", FirstWeakAttack: "頭部", SecondWeakAttack: "翼", FirstWeakElement: "龍", SecondWeakElement: "雷"},
	}
	weak2 := []*Weakness{
		{MonsterId: "0000000002", PartId: "0001", Fire: "45", Water: "45", Lightning: "45", Ice: "45", Dragon: "45", Slashing: "45", Blow: "45", Bullet: "45", FirstWeakAttack: "頭部", SecondWeakAttack: "翼", FirstWeakElement: "龍", SecondWeakElement: "雷"},
	}
	weak3 := []*Weakness{
		{MonsterId: "0000000003", PartId: "0001", Fire: "45", Water: "45", Lightning: "45", Ice: "45", Dragon: "45", Slashing: "45", Blow: "45", Bullet: "45", FirstWeakAttack: "頭部", SecondWeakAttack: "翼", FirstWeakElement: "雷", SecondWeakElement: "水"},
	}
	monsters := []Monster{
		{MonsterId: "0000000001", Name: "リオレウス", Description: "空の王者。", Field: []*Field{{FieldId: "0001", MonsterId: "0000000001", Name: "古代樹の森", ImageUrl: "images/kodaizyu.png"}}, Tribe: &Tribe{TribeId: "0001", MonsterId: "0000000001", Name_ja: "飛竜種", Name_en: "wibarn", Description: "飛竜種"}, Product: []*Product{{ProductId: "0001", MonsterId: "0000000001", Name: "MH", PublishYear: "2004", TotalSales: "200万本"}}, Weakness: weak1, Ranking: []*Ranking{{MonsterId: "0000000001", Ranking: "1", VoteYear: "2024/3/12"}}},
		{MonsterId: "0000000002", Name: "リオレイア", Description: "陸の女王", Field: []*Field{{FieldId: "0001", MonsterId: "0000000002", Name: "古代樹の森", ImageUrl: "images/kodaizyu.png"}}, Tribe: &Tribe{TribeId: "0001", MonsterId: "0000000002", Name_ja: "飛竜種", Name_en: "wibarn", Description: "飛竜種"}, Product: []*Product{{ProductId: "0001", MonsterId: "0000000002", Name: "MH", PublishYear: "2004", TotalSales: "200万本"}}, Weakness: weak2, Ranking: []*Ranking{{MonsterId: "0000000002", Ranking: "2", VoteYear: "2024/3/12"}}},
		{MonsterId: "0000000003", Name: "ティガレックス", Description: "絶対強者", Field: []*Field{{FieldId: "0001", MonsterId: "0000000003", Name: "古代樹の森", ImageUrl: "images/kodaizyu.png"}}, Tribe: &Tribe{TribeId: "0001", MonsterId: "0000000003", Name_ja: "飛竜種", Name_en: "wibarn", Description: "飛竜種"}, Product: []*Product{{ProductId: "0001", MonsterId: "0000000003", Name: "MH", PublishYear: "2004", TotalSales: "200万本"}}, Weakness: weak3, Ranking: []*Ranking{{MonsterId: "0000000003", Ranking: "3", VoteYear: "2024/3/12"}}},
	}
	bgmRanks1 := []*BgmRanking{
		{MusicId: "0000000001", Ranking: "1", VoteYear: "2024/3/12"},
	}

	bgmRanks2 := []*BgmRanking{
		{MusicId: "0000000002", Ranking: "2", VoteYear: "2024/3/12"},
	}

	bgmRanks3 := []*BgmRanking{
		{MusicId: "0000000003", Ranking: "3", VoteYear: "2024/3/12"},
	}

	bgms := []Music{
		{MusicId: "0000000001", MonsterId: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1", BgmRanking: bgmRanks1},
		{MusicId: "0000000002", MonsterId: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2", BgmRanking: bgmRanks2},
		{MusicId: "0000000003", MonsterId: "0000000003", Name: "ティガレックスのテーマ", Url: "https://www.youtube.com/watch?v=3", BgmRanking: bgmRanks3},
	}

	db.Exec("SET foreign_key_checks = 0")
	db.Exec("TRUNCATE TABLE monster")
	db.Exec("TRUNCATE TABLE field")
	db.Exec("TRUNCATE TABLE product")
	db.Exec("TRUNCATE TABLE tribe")
	db.Exec("TRUNCATE TABLE weakness")
	db.Exec("TRUNCATE TABLE ranking")
	db.Exec("TRUNCATE TABLE music")
	db.Exec("TRUNCATE TABLE bgm_ranking")
	db.Exec("SET foreign_key_checks = 1")
	db.Create(monsters)
	db.Create(bgms)
}

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

// cleanupTestDB はテスト用のDBクリーンアップを行うヘルパー関数
func cleanupTestDB(_ context.Context) {
	// TRUNCATE文のSQLの読み込み
	truncateSQL, err := os.ReadFile("../../../db/seed/00_trancate.sql")
	if err != nil {
		log.Fatal(err)
	}

	// テーブルのクリーンアップ
	statements := strings.Split(string(truncateSQL), ";")
	// テーブルのクリーンアップ
	for _, stmt := range statements {
		// 空の文を除外
		if strings.TrimSpace(stmt) != "" {
			if err := testDB.Exec(stmt).Error; err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Println("successfully cleaned up tables!")
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
	statements := strings.Split(string(truncateSQL), ";")
	// テーブルのクリーンアップ
	for _, stmt := range statements {
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
