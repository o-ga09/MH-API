package mysql

import (
	param "mh-api/internal/controller/monster"
	"mh-api/internal/domain/music"
	"mh-api/internal/service/monsters"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func Test_monsterQueryService_FetchList(t *testing.T) {
	ctx := t.Context()
	ctx = setupTestDB(ctx)
	db = ctx.Value(CtxKey).(*gorm.DB)
	db.Begin()
	defer db.Rollback()

	_ = createMonsterData(t, ctx)

	weak_A := []monsters.Weakness_attack{
		{PartId: "0001", Slashing: "45", Blow: "45", Bullet: "45"},
	}
	weak_E := []monsters.Weakness_element{
		{PartId: "0001", Fire: "45", Water: "45", Thunder: "45", Ice: "45", Dragon: "45"},
	}

	// ランキングとBGM情報を追加
	ranking1 := []monsters.Ranking{
		{Ranking: "1", VoteYear: "2024/3/12"},
	}
	ranking2 := []monsters.Ranking{
		{Ranking: "2", VoteYear: "2024/3/12"},
	}
	ranking3 := []monsters.Ranking{
		{Ranking: "3", VoteYear: "2024/3/12"},
	}

	bgm1 := music.NewMusic("0000000001", "0000000001", "リオレウスのテーマ", "https://www.youtube.com/watch?v=1")
	bgm2 := music.NewMusic("0000000002", "0000000002", "リオレイアのテーマ", "https://www.youtube.com/watch?v=2")
	bgm3 := music.NewMusic("0000000003", "0000000003", "ティガレックスのテーマ", "https://www.youtube.com/watch?v=3")

	// ID順に並んでいるデータを適応
	monster1 := &monsters.FetchMonsterListDto{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E, Ranking: ranking1, BGM: []music.Music{*bgm1}}
	monster2 := &monsters.FetchMonsterListDto{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E, Ranking: ranking2, BGM: []music.Music{*bgm2}}
	monster3 := &monsters.FetchMonsterListDto{Id: "0000000003", Name: "ティガレックス", Description: "絶対強者", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "雷", SecondWeak_Element: "水", Weakness_attack: weak_A, Weakness_element: weak_E, Ranking: ranking3, BGM: []music.Music{*bgm3}}

	param1 := param.RequestParam{MonsterIds: "", MonsterName: "", Limit: 100, Offset: 0, Sort: "1"}
	param2 := param.RequestParam{MonsterIds: "0000000001,0000000002", MonsterName: "", Limit: 100, Offset: 0, Sort: "1"}
	param3 := param.RequestParam{MonsterIds: "", MonsterName: "リオレウス", Limit: 100, Offset: 0, Sort: "1"}
	param4 := param.RequestParam{MonsterIds: "", MonsterName: "", Limit: 100, Offset: 0, Sort: "1"}
	param5 := param.RequestParam{MonsterIds: "", MonsterName: "", Limit: 100, Offset: 0, Sort: "2"}
	param6 := param.RequestParam{MonsterIds: "0000000001", MonsterName: "リオレウス", Limit: 100, Offset: 0, Sort: "1"}

	type args struct {
		id    string
		param param.RequestParam
	}
	tests := []struct {
		name      string
		args      args
		want      []*monsters.FetchMonsterListDto
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "DBからモンスターデータを複数件取得できる",
			args:      args{id: "", param: param1},
			want:      []*monsters.FetchMonsterListDto{monster3, monster2, monster1},
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name:      "DBからモンスターデータをmonsterIdを複数件指定して取得できる",
			args:      args{id: "", param: param2},
			want:      []*monsters.FetchMonsterListDto{monster2, monster1},
			wantTotal: 2,
			wantErr:   false,
		},
		{
			name:      "DBからモンスターの名前を部分一致検索で指定して取得できる",
			args:      args{id: "", param: param3},
			want:      []*monsters.FetchMonsterListDto{monster1},
			wantTotal: 1,
			wantErr:   false,
		},
		{
			name:      "DBからモンスターデータをmonsterIdでソート（昇順）して取得できる",
			args:      args{id: "", param: param4},
			want:      []*monsters.FetchMonsterListDto{monster3, monster2, monster1},
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name:      "DBからモンスターデータをmonsterIdでソート（降順）して取得できる",
			args:      args{id: "", param: param5},
			want:      []*monsters.FetchMonsterListDto{monster1, monster2, monster3},
			wantTotal: 3,
			wantErr:   false,
		},
		{
			name:      "DBからモンスターデータをid指定で1件取得できる",
			args:      args{id: "0000000002", param: param.RequestParam{}},
			want:      []*monsters.FetchMonsterListDto{monster2},
			wantTotal: 1,
			wantErr:   false,
		},
		{
			name:      "DBからモンスターデータを取得できない場合、NotFoundErrorで返す",
			args:      args{id: "", param: param.RequestParam{}},
			want:      nil,
			wantTotal: 0,
			wantErr:   true},
		{
			name:      "DBからモンスターデータをmonsterIdとmonsterNameを指定して取得できる",
			args:      args{id: "", param: param6},
			want:      []*monsters.FetchMonsterListDto{monster1},
			wantTotal: 1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx = context.WithValue(ctx, "param", tt.args.param)
			s := &monsterQueryService{}
			got, err := s.FetchList(ctx, tt.args.id)
			assert.True(t, (err != nil) == tt.wantErr)

			if !tt.wantErr {
				require.NotNil(t, got)
				assert.Equal(t, tt.wantTotal, got.Total)
				assert.Equal(t, len(tt.want), len(got.Monsters))
				assert.Equal(t, tt.want, got.Monsters)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func createMonsterData(t *testing.T, ctx context.Context) []*monsters.FetchMonsterListDto {
	t.Helper()

	weak1 := []*Weakness{
		{MonsterId: "0000000001", PartId: "0001", Fire: "45", Water: "45", Lightning: "45", Ice: "45", Dragon: "45", Slashing: "45", Blow: "45", Bullet: "45", FirstWeakAttack: "頭部", SecondWeakAttack: "翼", FirstWeakElement: "龍", SecondWeakElement: "雷"},
	}
	weak2 := []*Weakness{
		{MonsterId: "0000000002", PartId: "0001", Fire: "45", Water: "45", Lightning: "45", Ice: "45", Dragon: "45", Slashing: "45", Blow: "45", Bullet: "45", FirstWeakAttack: "頭部", SecondWeakAttack: "翼", FirstWeakElement: "龍", SecondWeakElement: "雷"},
	}
	weak3 := []*Weakness{
		{MonsterId: "0000000003", PartId: "0001", Fire: "45", Water: "45", Lightning: "45", Ice: "45", Dragon: "45", Slashing: "45", Blow: "45", Bullet: "45", FirstWeakAttack: "頭部", SecondWeakAttack: "翼", FirstWeakElement: "雷", SecondWeakElement: "水"},
	}
	monster := []Monster{
		{MonsterId: "0000000001", Name: "リオレウス", Description: "空の王者。", Field: []*Field{{FieldId: "0001", MonsterId: "0000000001", Name: "古代樹の森", ImageUrl: "images/kodaizyu.png"}}, Tribe: &Tribe{TribeId: "0001", MonsterId: "0000000001", Name_ja: "飛竜種", Name_en: "wibarn", Description: "飛竜種"}, Product: []*Product{{ProductId: "0001", MonsterId: "0000000001", Name: "MH", PublishYear: "2004", TotalSales: "200万本"}}, Weakness: weak1, Ranking: []*Ranking{{MonsterId: "0000000001", Ranking: "1", VoteYear: "2024/3/12"}}},
		{MonsterId: "0000000002", Name: "リオレイア", Description: "陸の女王", Field: []*Field{{FieldId: "0001", MonsterId: "0000000002", Name: "古代樹の森", ImageUrl: "images/kodaizyu.png"}}, Tribe: &Tribe{TribeId: "0001", MonsterId: "0000000002", Name_ja: "飛竜種", Name_en: "wibarn", Description: "飛竜種"}, Product: []*Product{{ProductId: "0001", MonsterId: "0000000002", Name: "MH", PublishYear: "2004", TotalSales: "200万本"}}, Weakness: weak2, Ranking: []*Ranking{{MonsterId: "0000000002", Ranking: "2", VoteYear: "2024/3/12"}}},
		{MonsterId: "0000000003", Name: "ティガレックス", Description: "絶対強者", Field: []*Field{{FieldId: "0001", MonsterId: "0000000003", Name: "古代樹の森", ImageUrl: "images/kodaizyu.png"}}, Tribe: &Tribe{TribeId: "0001", MonsterId: "0000000003", Name_ja: "飛竜種", Name_en: "wibarn", Description: "飛竜種"}, Product: []*Product{{ProductId: "0001", MonsterId: "0000000003", Name: "MH", PublishYear: "2004", TotalSales: "200万本"}}, Weakness: weak3, Ranking: []*Ranking{{MonsterId: "0000000003", Ranking: "3", VoteYear: "2024/3/12"}}},
	}

	bgms := []Music{
		{MusicId: "0000000001", MonsterId: "0000000001", Name: "リオレウスのテーマ", Url: "https://www.youtube.com/watch?v=1"},
		{MusicId: "0000000002", MonsterId: "0000000002", Name: "リオレイアのテーマ", Url: "https://www.youtube.com/watch?v=2"},
		{MusicId: "0000000003", MonsterId: "0000000003", Name: "ティガレックスのテーマ", Url: "https://www.youtube.com/watch?v=3"},
	}

	err := CtxFromTestDB(ctx).Create(monster).Error
	require.NoError(t, err)

	err = CtxFromTestDB(ctx).Create(bgms).Error
	require.NoError(t, err)

	result := []*monsters.FetchMonsterListDto{}
	for _, m := range monster {
		dto := MonsterToDTO(m)
		result = append(result, dto)
	}

	return result
}
