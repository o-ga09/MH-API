package mysql

import (
	"context"
	"log/slog"
	"mh-api/app/internal/presenter/middleware"
	"mh-api/app/pkg"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func BeforeTest() {
	var err error
	os.Setenv("DATABASE_URL", "root:pass@tcp(127.0.0.1:3306)/ci?charset=utf8&parseTime=True&loc=Local")

	cfg, err := pkg.New()
	if err != nil {
		slog.Log(context.Background(), middleware.SeverityError, "environment variable error", "error", err)
	}
	dialector := mysql.Open(cfg.Database_url)

	var db *gorm.DB

	if db, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		connect(dialector, 100)
	}
	err = db.AutoMigrate(&Monster{}, &Field{}, &Product{}, &Tribe{}, &Weakness{}, &Ranking{}, &Music{}, &BgmRanking{}, &Item{}, &ItemWithMonster{})
	if err != nil {
		panic(err)
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

	items := []Item{
		{ItemId: "0000000001", Name: "回復薬", NameKana: "カイフクヤク", ImageUrl: "images/rioreusu.png"},
		{ItemId: "0000000002", Name: "回復薬グレート", NameKana: "カイフクヤクグレート", ImageUrl: "images/rioreia.png"},
		{ItemId: "0000000003", Name: "秘薬", NameKana: "ヒヤク", ImageUrl: "images/tigarekkusu.png"},
		{ItemId: "0000000004", Name: "砥石", NameKana: "トイシ", ImageUrl: "images/tigarekkusu.png"},
		{ItemId: "0000000005", Name: "おとし穴", NameKana: "オトシアナ", ImageUrl: "images/tigarekkusu.png"},
		{ItemId: "0000000006", Name: "毒ビン", NameKana: "ドクビン", ImageUrl: "images/tigarekkusu.png"},
		{ItemId: "0000000007", Name: "麻痺ビン", NameKana: "マヒビン", ImageUrl: "images/tigarekkusu.png"},
		{ItemId: "0000000008", Name: "眠りビン", NameKana: "ネムリビン", ImageUrl: "images/tigarekkusu.png"},
		{ItemId: "0000000009", Name: "爆弾", NameKana: "バクダン", ImageUrl: "images/tigarekkusu.png"},
		{ItemId: "0000000010", Name: "大タル爆弾", NameKana: "オオタルバクダン", ImageUrl: "images/tigarekkusu.png"},
		{ItemId: "0000000011", Name: "閃光玉", NameKana: "センコウダマ", ImageUrl: "images/tigarekkusu.png"},
	}

	itemsWithMonster := []ItemWithMonster{
		{ItemId: "0000000001", MonsterId: "0000000001"},
		{ItemId: "0000000001", MonsterId: "0000000002"},
		{ItemId: "0000000001", MonsterId: "0000000003"},
		{ItemId: "0000000002", MonsterId: "0000000001"},
		{ItemId: "0000000002", MonsterId: "0000000002"},
		{ItemId: "0000000002", MonsterId: "0000000003"},
		{ItemId: "0000000003", MonsterId: "0000000001"},
		{ItemId: "0000000003", MonsterId: "0000000002"},
		{ItemId: "0000000003", MonsterId: "0000000003"},
	}
	db.Create(monsters)
	db.Create(bgms)
	db.Create(items)
	db.Create(itemsWithMonster)
}

func AfetrTest() func() {
	return func() {
		var err error

		cfg, err := pkg.New()
		if err != nil {
			slog.Log(context.Background(), middleware.SeverityError, "environment variable error", "error", err)
		}
		dialector := mysql.Open(cfg.Database_url)

		var db *gorm.DB

		if db, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}}); err != nil {
			connect(dialector, 100)
		}
		db.Exec("SET foreign_key_checks = 0")
		db.Exec("DROP TABLE monster")
		db.Exec("DROP TABLE field")
		db.Exec("DROP TABLE product")
		db.Exec("DROP TABLE tribe")
		db.Exec("DROP TABLE weakness")
		db.Exec("DROP TABLE ranking")
		db.Exec("DROP TABLE music")
		db.Exec("DROP TABLE bgm_ranking")
		db.Exec("DROP TABLE item")
		db.Exec("DROP TABLE item_with_monster")
		db.Exec("SET foreign_key_checks = 1")

		// db.Exec("DELETE FROM monster")
		// db.Exec("DELETE FROM field")
		// db.Exec("DELETE FROM product")
		// db.Exec("DELETE FROM tribe")
		// db.Exec("DELETE FROM weakness")
		// db.Exec("DELETE FROM ranking")
		// db.Exec("DELETE FROM music")
		// db.Exec("DELETE FROM bgm_ranking")
		// db.Exec("DELETE FROM item")
		// db.Exec("DELETE FROM item_with_monster")
		// db.Exec("DROP INDEX idx_monster_deleted_at ON monster")
		// db.Exec("DROP INDEX idx_monster_monster_id ON monster")
		// db.Exec("DROP INDEX idx_field_deleted_at ON field")
		// db.Exec("DROP INDEX idx_field_field_id ON field")
		// db.Exec("DROP INDEX idx_product_deleted_at ON product")
		// db.Exec("DROP INDEX idx_product_product_id ON product")
		// db.Exec("DROP INDEX idx_tribe_deleted_at ON tribe")
		// db.Exec("DROP INDEX idx_tribe_tribe_id ON tribe")
		// db.Exec("DROP INDEX idx_weakness_deleted_at ON weakness")
		// db.Exec("DROP INDEX idx_weakness_weakness_id ON weakness")
		// db.Exec("DROP INDEX idx_ranking_deleted_at ON ranking")
		// db.Exec("DROP INDEX idx_ranking_ranking_id ON ranking")
		// db.Exec("DROP INDEX idx_music_deleted_at ON music")
		// db.Exec("DROP INDEX idx_music_music_id ON music")
		// db.Exec("DROP INDEX idx_bgm_ranking_deleted_at ON bgm_ranking")
		// db.Exec("DROP INDEX idx_bgm_ranking_bgm_ranking_id ON bgm_ranking")
		// db.Exec("DROP INDEX idx_item_deleted_at ON item")
		// db.Exec("DROP INDEX idx_item_item_id ON item")
		// db.Exec("DROP INDEX idx_item_with_monster_deleted_at ON item_with_monster")
		// db.Exec("DROP INDEX idx_item_with_monster_item_with_monster_id ON item_with_monster")
		// db.Exec("ALTER TABLE monster AUTO_INCREMENT = 1")
		// db.Exec("ALTER TABLE field AUTO_INCREMENT = 1")
		// db.Exec("ALTER TABLE product AUTO_INCREMENT = 1")
		// db.Exec("ALTER TABLE tribe AUTO_INCREMENT = 1")
		// db.Exec("ALTER TABLE weakness AUTO_INCREMENT = 1")
		// db.Exec("ALTER TABLE ranking AUTO_INCREMENT = 1")
		// db.Exec("ALTER TABLE music AUTO_INCREMENT = 1")
		// db.Exec("ALTER TABLE bgm_ranking AUTO_INCREMENT = 1")
		// db.Exec("ALTER TABLE item AUTO_INCREMENT = 1")
		// db.Exec("ALTER TABLE item_with_monster AUTO_INCREMENT = 1")
		// db.Exec("SET foreign_key_checks = 1")
	}
}
