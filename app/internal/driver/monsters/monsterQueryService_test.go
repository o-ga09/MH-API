package monsters

import (
	param "mh-api/app/internal/controller/monster"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/monsters"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/context"
)

func TestNewmonsterQueryService(t *testing.T) {
	t.Skip()
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	tests := []struct {
		name string
		want *monsterQueryService
	}{
		{name: "QueryService構造体を生成する", want: NewmonsterQueryService()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewmonsterQueryService()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewmonsterQueryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monsterQueryService_FetchList(t *testing.T) {
	t.Skip()
	t.Helper()
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	ctx := context.Background()

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
	wantMonsters7 := []*monsters.FetchMonsterListDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Location: []string{"古代樹の森"}, Category: "飛竜種", Title: []string{"MH"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
	}
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
		name    string
		args    args
		want    []*monsters.FetchMonsterListDto
		wantErr bool
	}{
		{name: "DBからモンスターデータを複数件取得できる", args: args{id: "", param: param1}, want: wantMonsters1, wantErr: false},
		{name: "DBからモンスターデータをmonsterIdを複数件指定して取得できる", args: args{id: "", param: param2}, want: wantMonsters2, wantErr: false},
		{name: "DBからモンスターの名前を部分一致検索で指定して取得できる", args: args{id: "", param: param3}, want: wantMonsters3, wantErr: false},
		{name: "DBからモンスターデータをmonsterIdでソート（昇順）して取得できる", args: args{id: "", param: param4}, want: wantMonsters4, wantErr: false},
		{name: "DBからモンスターデータをmonsterIdでソート（降順）して取得できる", args: args{id: "", param: param5}, want: wantMonsters5, wantErr: false},
		{name: "DBからモンスターデータをid指定で1件取得できる", args: args{id: "0000000002", param: param.RequestParam{}}, want: wantMonsters6, wantErr: false},
		{name: "DBからモンスターデータを取得できない場合、NotFoundErrorで返す", args: args{id: "", param: param.RequestParam{}}, want: nil, wantErr: true},
		{name: "DBからモンスターデータをmonsterIdとmonsterNameを指定して取得できる", args: args{id: "", param: param6}, want: wantMonsters7, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx = context.WithValue(context.Background(), "param", tt.args.param)
			s := &monsterQueryService{}
			got, err := s.FetchList(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("monsterQueryService.FetchList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); !tt.wantErr && diff != "" {
				t.Errorf("monsterQueryService.FetchList() = %v, want %v", got, tt.want)
			}
		})
	}
}
