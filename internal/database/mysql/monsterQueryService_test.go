package mysql

import (
	"context"
	"testing"

	"mh-api/internal/domain/fields"
	"mh-api/internal/domain/monsters"
	"mh-api/internal/domain/music"
	Products "mh-api/internal/domain/products"
	"mh-api/internal/domain/ranking"
	Tribes "mh-api/internal/domain/tribes"
	"mh-api/internal/domain/weakness"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func Test_monsterRepository_FindAll(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	db.Begin()
	defer db.Rollback()

	createMonsterData(t, ctx)

	noTribeMonster := &monsters.Monster{MonsterId: "0000000004", Name: "NoTribeMonster", Description: "This monster has no tribe."}
	require.NoError(t, CtxFromTestDB(ctx).Create(noTribeMonster).Error)

	tests := []struct {
		name      string
		params    monsters.SearchParams
		wantIDs   []string
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "全件取得（降順）",
			params:    monsters.SearchParams{Limit: 100, Sort: "desc"},
			wantIDs:   []string{"0000000004", "0000000003", "0000000002", "0000000001"},
			wantTotal: 4,
		},
		{
			name:      "monsterIdを複数指定",
			params:    monsters.SearchParams{MonsterIds: "0000000001,0000000002", Limit: 100, Sort: "desc"},
			wantIDs:   []string{"0000000002", "0000000001"},
			wantTotal: 2,
		},
		{
			name:      "名前部分一致検索",
			params:    monsters.SearchParams{MonsterName: "リオレウス", Limit: 100},
			wantIDs:   []string{"0000000001"},
			wantTotal: 1,
		},
		{
			name:      "昇順",
			params:    monsters.SearchParams{Limit: 100},
			wantIDs:   []string{"0000000001", "0000000002", "0000000003", "0000000004"},
			wantTotal: 4,
		},
		{
			name:      "降順（Sort=1）",
			params:    monsters.SearchParams{Limit: 100, Sort: "desc"},
			wantIDs:   []string{"0000000004", "0000000003", "0000000002", "0000000001"},
			wantTotal: 4,
		},
		{
			name:      "使用属性Fireで検索",
			params:    monsters.SearchParams{UsageElement: "Fire", Limit: 100},
			wantIDs:   []string{"0000000001"},
			wantTotal: 1,
		},
		{
			name:      "弱点属性「龍」で検索",
			params:    monsters.SearchParams{WeaknessElement: "龍", Limit: 100},
			wantIDs:   []string{"0000000001", "0000000002"},
			wantTotal: 2,
		},
		{
			name:      "弱点属性「Dragon」（英語）で検索",
			params:    monsters.SearchParams{WeaknessElement: "Dragon", Limit: 100},
			wantIDs:   []string{"0000000001", "0000000002"},
			wantTotal: 2,
		},
		{
			name:      "使用属性Water＋弱点属性「龍」の複合検索",
			params:    monsters.SearchParams{UsageElement: "Water", WeaknessElement: "龍", Limit: 100},
			wantIDs:   []string{"0000000002"},
			wantTotal: 1,
		},
		{
			name:      "Tribeなしモンスターでpanicしない",
			params:    monsters.SearchParams{MonsterIds: "0000000004", Limit: 100},
			wantIDs:   []string{"0000000004"},
			wantTotal: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &monsterRepository{}
			got, err := repo.FindAll(ctx, tt.params)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, got)
			assert.Equal(t, tt.wantTotal, got.Total)
			require.Len(t, got.Monsters, len(tt.wantIDs))
			for i, id := range tt.wantIDs {
				assert.Equal(t, id, got.Monsters[i].MonsterId)
			}
		})
	}
}

func Test_monsterRepository_FindById(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	db.Begin()
	defer db.Rollback()

	createMonsterData(t, ctx)

	tests := []struct {
		name    string
		id      string
		wantID  string
		wantErr bool
	}{
		{name: "正常系: 存在するID", id: "0000000001", wantID: "0000000001"},
		{name: "異常系: 存在しないID", id: "9999999999", wantErr: true},
		{name: "異常系: 空のID", id: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &monsterRepository{}
			got, err := repo.FindById(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantID, got.MonsterId)
			assert.NotEmpty(t, got.Name)
		})
	}
}

func createMonsterData(t *testing.T, ctx context.Context) {
	t.Helper()

	fireElement := "火"
	waterElement := "水"
	dragonElement := "龍"

	monsterList := []*monsters.Monster{
		{MonsterId: "0000000001", Name: "リオレウス", Description: "空の王者。", Element: &fireElement},
		{MonsterId: "0000000002", Name: "リオレイア", Description: "陸の女王", Element: &waterElement},
		{MonsterId: "0000000003", Name: "ティガレックス", Description: "絶対強者", Element: &dragonElement},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(monsterList).Error)

	weaknesses := []*weakness.Weakness{
		{MonsterId: "0000000001", PartId: "0001", Fire: "45", Water: "45", Lightning: "45", Ice: "45", Dragon: "45", Slashing: "45", Blow: "45", Bullet: "45", FirstWeakAttack: "頭部", SecondWeakAttack: "翼", FirstWeakElement: "龍", SecondWeakElement: "雷"},
		{MonsterId: "0000000002", PartId: "0001", Fire: "45", Water: "45", Lightning: "45", Ice: "45", Dragon: "45", Slashing: "45", Blow: "45", Bullet: "45", FirstWeakAttack: "頭部", SecondWeakAttack: "翼", FirstWeakElement: "龍", SecondWeakElement: "雷"},
		{MonsterId: "0000000003", PartId: "0001", Fire: "45", Water: "45", Lightning: "45", Ice: "45", Dragon: "45", Slashing: "45", Blow: "45", Bullet: "45", FirstWeakAttack: "頭部", SecondWeakAttack: "翼", FirstWeakElement: "雷", SecondWeakElement: "水"},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(weaknesses).Error)

	fieldList := []*fields.Field{
		{FieldId: "0001", MonsterId: "0000000001", Name: "古代樹の森", ImageUrl: "images/kodaizyu.png"},
		{FieldId: "0002", MonsterId: "0000000002", Name: "古代樹の森", ImageUrl: "images/kodaizyu.png"},
		{FieldId: "0003", MonsterId: "0000000003", Name: "古代樹の森", ImageUrl: "images/kodaizyu.png"},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(fieldList).Error)

	tribeList := []*Tribes.Tribe{
		{TribeId: "0001", MonsterId: "0000000001", Name_ja: "飛竜種", Name_en: "wibarn", Description: "飛竜種"},
		{TribeId: "0002", MonsterId: "0000000002", Name_ja: "飛竜種", Name_en: "wibarn", Description: "飛竜種"},
		{TribeId: "0003", MonsterId: "0000000003", Name_ja: "飛竜種", Name_en: "wibarn", Description: "飛竜種"},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(tribeList).Error)

	productList := []*Products.Product{
		{ProductId: "0001", MonsterId: "0000000001", Name: "MH", PublishYear: "2004", TotalSales: "200万本"},
		{ProductId: "0002", MonsterId: "0000000002", Name: "MH", PublishYear: "2004", TotalSales: "200万本"},
		{ProductId: "0003", MonsterId: "0000000003", Name: "MH", PublishYear: "2004", TotalSales: "200万本"},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(productList).Error)

	rankingList := []*ranking.Ranking{
		{MonsterId: "0000000001", Ranking: "1", VoteYear: "2024/3/12"},
		{MonsterId: "0000000002", Ranking: "2", VoteYear: "2024/3/12"},
		{MonsterId: "0000000003", Ranking: "3", VoteYear: "2024/3/12"},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(rankingList).Error)

	bgmList := []*music.Music{
		{MusicId: "0000000001", MonsterId: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1"},
		{MusicId: "0000000002", MonsterId: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2"},
		{MusicId: "0000000003", MonsterId: "0000000003", Name: "ティガレックスのテーマ", Url: "https://www.youtube.com/watch?v=3"},
	}
	require.NoError(t, CtxFromTestDB(ctx).Create(bgmList).Error)
}
