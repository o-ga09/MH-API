package mysql

import (
	"fmt"
	param "mh-api/app/internal/controller/monster"
	"mh-api/app/internal/service/monsters"
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/context"

	"gorm.io/gorm"
)

func TestNewmonsterQueryService(t *testing.T) {
	os.Setenv("DATABASE_URL", "mh-api:P@ssw0rd@tcp(127.0.0.1:3306)/mh-api?charset=utf8&parseTime=True&loc=Local")
	conn := New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *monsterQueryService
	}{
		{name: "QueryService構造体を生成する", args: args{conn: conn}, want: NewmonsterQueryService(conn)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewmonsterQueryService(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewmonsterQueryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monsterQueryService_FetchList(t *testing.T) {
	t.Helper()
	BeforeTest()
	t.Cleanup(AfetrTest())
	ctx := context.Background()
	conn := New(ctx)
	weak_A := []monsters.Weakness_attack{
		{PartId: "0001", Slashing: "45", Blow: "45", Bullet: "45"},
	}
	weak_E := []monsters.Weakness_element{
		{PartId: "0001", Fire: "45", Water: "45", Thunder: "45", Ice: "45", Dragon: "45"},
	}
	wantMonsters1 := []*monsters.FetchMonsterListDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
		{Id: "0000000003", Name: "ティガレックス", Description: "絶対強者", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "雷", SecondWeak_Element: "水", Weakness_attack: weak_A, Weakness_element: weak_E},
	}
	wantMonsters2 := []*monsters.FetchMonsterListDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
	}
	wantMonsters3 := []*monsters.FetchMonsterListDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
	}
	wantMonsters4 := []*monsters.FetchMonsterListDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
		{Id: "0000000003", Name: "ティガレックス", Description: "絶対強者", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "雷", SecondWeak_Element: "水", Weakness_attack: weak_A, Weakness_element: weak_E},
	}
	wantMonsters5 := []*monsters.FetchMonsterListDto{
		{Id: "0000000003", Name: "ティガレックス", Description: "絶対強者", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "雷", SecondWeak_Element: "水", Weakness_attack: weak_A, Weakness_element: weak_E},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
	}
	wantMonsters6 := []*monsters.FetchMonsterListDto{
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
	}
	param1 := param.RequestParam{MonsterIds: "", MonsterName: "", Limit: 100, Offset: 0, Sort: "1"}
	param2 := param.RequestParam{MonsterIds: "0000000001,0000000002", MonsterName: "", Limit: 100, Offset: 0, Sort: "1"}
	param3 := param.RequestParam{MonsterIds: "", MonsterName: "リオレウス", Limit: 100, Offset: 0, Sort: "1"}
	param4 := param.RequestParam{MonsterIds: "", MonsterName: "", Limit: 100, Offset: 0, Sort: "1"}
	param5 := param.RequestParam{MonsterIds: "", MonsterName: "", Limit: 100, Offset: 0, Sort: "2"}

	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		id    string
		param param.RequestParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*monsters.FetchMonsterListDto
		wantErr bool
	}{
		{name: "DBからモンスターデータを複数件取得できる", fields: fields{conn: conn}, args: args{id: "", param: param1}, want: wantMonsters1, wantErr: false},
		{name: "DBからモンスターデータをmonsterIdを複数件指定して取得できる", fields: fields{conn: conn}, args: args{id: "", param: param2}, want: wantMonsters2, wantErr: false},
		{name: "DBからモンスターの名前を部分一致検索で指定して取得できる", fields: fields{conn: conn}, args: args{id: "", param: param3}, want: wantMonsters3, wantErr: false},
		{name: "DBからモンスターデータをmonsterIdでソート（昇順）して取得できる", fields: fields{conn: conn}, args: args{id: "", param: param4}, want: wantMonsters4, wantErr: false},
		{name: "DBからモンスターデータをmonsterIdでソート（降順）して取得できる", fields: fields{conn: conn}, args: args{id: "", param: param5}, want: wantMonsters5, wantErr: false},
		{name: "DBからモンスターデータをid指定で1件取得できる", fields: fields{conn: conn}, args: args{id: "0000000002", param: param.RequestParam{}}, want: wantMonsters6, wantErr: false},
		{name: "DBからモンスターデータを取得できない場合、NotFoundErrorで返す", fields: fields{conn: conn}, args: args{id: "", param: param.RequestParam{}}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx = context.WithValue(context.Background(), "param", tt.args.param)
			s := &monsterQueryService{
				conn: tt.fields.conn,
			}
			got, err := s.FetchList(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("monsterQueryService.FetchList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); !tt.wantErr && diff != "" {
				fmt.Println("got : ", got[0])
				fmt.Println("want : ", tt.want[0])
				t.Errorf("monsterQueryService.FetchList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monsterQueryService_FetchRank(t *testing.T) {
	t.Helper()
	BeforeTest()
	t.Cleanup(AfetrTest())
	ctx := context.Background()
	conn := New(ctx)
	wantMonsters1 := []*monsters.FetchMonsterRankingDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "1", VoteYear: "2024/3/12"}}},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "2", VoteYear: "2024/3/12"}}},
		{Id: "0000000003", Name: "ティガレックス", Description: "絶対強者", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "3", VoteYear: "2024/3/12"}}},
	}
	wantMonsters2 := []*monsters.FetchMonsterRankingDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "1", VoteYear: "2024/3/12"}}},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "2", VoteYear: "2024/3/12"}}},
	}
	wantMonsters3 := []*monsters.FetchMonsterRankingDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "1", VoteYear: "2024/3/12"}}},
	}
	wantMonsters4 := []*monsters.FetchMonsterRankingDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "1", VoteYear: "2024/3/12"}}},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "2", VoteYear: "2024/3/12"}}},
		{Id: "0000000003", Name: "ティガレックス", Description: "絶対強者", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "3", VoteYear: "2024/3/12"}}},
	}
	wantMonsters5 := []*monsters.FetchMonsterRankingDto{
		{Id: "0000000003", Name: "ティガレックス", Description: "絶対強者", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "3", VoteYear: "2024/3/12"}}},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "2", VoteYear: "2024/3/12"}}},
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, Ranking: []monsters.Ranking{{Ranking: "1", VoteYear: "2024/3/12"}}},
	}
	param1 := param.RequestRankingParam{MonsterIds: "", MonsterName: "", LocationName: "", TribeName: "", Title: "", Limit: 10, Offset: 0, Sort: ""}
	param2 := param.RequestRankingParam{MonsterIds: "0000000001,0000000002", MonsterName: "", LocationName: "", TribeName: "", Title: "", Limit: 100, Offset: 0, Sort: "1"}
	param3 := param.RequestRankingParam{MonsterIds: "", MonsterName: "リオレウス", LocationName: "", TribeName: "", Title: "", Limit: 100, Offset: 0, Sort: "1"}
	param4 := param.RequestRankingParam{MonsterIds: "", MonsterName: "", LocationName: "", TribeName: "", Title: "", Limit: 100, Offset: 0, Sort: "1"}
	param5 := param.RequestRankingParam{MonsterIds: "", MonsterName: "", LocationName: "", TribeName: "", Title: "", Limit: 100, Offset: 0, Sort: "2"}
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		param param.RequestRankingParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*monsters.FetchMonsterRankingDto
		wantErr bool
	}{
		{name: "DBからモンスターの人気投票結果を複数件取得できる", fields: fields{conn: conn}, args: args{param: param1}, want: wantMonsters1, wantErr: false},
		{name: "DBからモンスタの人気投票結果をmonsterIdで複数件指定して取得できる", fields: fields{conn: conn}, args: args{param: param2}, want: wantMonsters2, wantErr: false},
		{name: "DBからモンスターの名前を部分一致検索で指定して取得できる", fields: fields{conn: conn}, args: args{param: param3}, want: wantMonsters3, wantErr: false},
		{name: "DBからモンスターの人気投票結果をmonsterIdでソート（昇順）して取得できる", fields: fields{conn: conn}, args: args{param: param4}, want: wantMonsters4, wantErr: false},
		{name: "DBからモンスターの人気投票結果をmonsterIdでソート（降順）して取得できる", fields: fields{conn: conn}, args: args{param: param5}, want: wantMonsters5, wantErr: false},
		{name: "DBからモンスターデータを取得できない場合、NotFoundErrorで返す", fields: fields{conn: conn}, args: args{param: param.RequestRankingParam{}}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx = context.WithValue(context.Background(), "param", tt.args.param)
			s := &monsterQueryService{
				conn: tt.fields.conn,
			}
			got, err := s.FetchRank(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("monsterQueryService.FetchRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("monsterQueryService.FetchRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
