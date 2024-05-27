package queryservice

import (
	"context"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/monsters"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewMonsterQueryService(t *testing.T) {
	mysql.BeforeTest()
	conn := mysql.New(context.Background())
	type args struct {
		qs *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want MonsterQueryService
	}{
		{name: "TestNewMonsterQueryService", args: args{qs: conn}, want: MonsterQueryService{qs: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterQueryService(tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterQueryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonsterQueryService_FetchList(t *testing.T) {
	want1 := []*monsters.FetchMonsterListDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Category: "飛竜種", Title: []string{"MH"}, Location: []string{"古代樹の森"}, Weakness_attack: []monsters.Weakness_attack{{PartId: "0001", PartName: "頭部", Slashing: "45", Blow: "45", Bullet: "45"}}, Weakness_element: []monsters.Weakness_element{{PartId: "0001", PartName: "頭部", Fire: "45", Water: "45", Thunder: "45", Ice: "45", Dragon: "45"}}},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Category: "飛竜種", Title: []string{"MH"}, Location: []string{"古代樹の森"}, Weakness_attack: []monsters.Weakness_attack{{PartId: "0001", PartName: "翼", Slashing: "45", Blow: "45", Bullet: "45"}}, Weakness_element: []monsters.Weakness_element{{PartId: "0001", PartName: "翼", Fire: "45", Water: "45", Thunder: "45", Ice: "45", Dragon: "45"}}},
		{Id: "0000000003", Name: "ティガレックス", Description: "絶対強者", Category: "飛竜種", Title: []string{"MH"}, Location: []string{"古代樹の森"}, Weakness_attack: []monsters.Weakness_attack{{PartId: "0001", PartName: "尻尾", Slashing: "45", Blow: "45", Bullet: "45"}}, Weakness_element: []monsters.Weakness_element{{PartId: "0001", PartName: "尻尾", Fire: "45", Water: "45", Thunder: "45", Ice: "45", Dragon: "45"}}},
	}
	want2 := []*monsters.FetchMonsterListDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Category: "飛竜種", Title: []string{"MH"}, Location: []string{"古代樹の森"}, Weakness_attack: []monsters.Weakness_attack{{PartId: "0001", PartName: "頭部", Slashing: "45", Blow: "45", Bullet: "45"}}, Weakness_element: []monsters.Weakness_element{{PartId: "0001", PartName: "頭部", Fire: "45", Water: "45", Thunder: "45", Ice: "45", Dragon: "45"}}},
	}
	mysql.BeforeTest()
	conn := mysql.New(context.Background())
	type fields struct {
		qs *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*monsters.FetchMonsterListDto
		wantErr bool
	}{
		{name: "TestMonsterQueryService_FetchList", fields: fields{qs: conn}, args: args{ctx: context.Background(), id: ""}, want: want1, wantErr: false},
		{name: "TestMonsterQueryService_FetchList", fields: fields{qs: conn}, args: args{ctx: context.Background(), id: "0000000001"}, want: want2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MonsterQueryService{
				qs: tt.fields.qs,
			}
			got, err := s.FetchList(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MonsterQueryService.FetchList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonsterQueryService.FetchList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonsterQueryService_FetchRank(t *testing.T) {
	wants := []*monsters.FetchMonsterRankingDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の王者。", Category: "飛竜種", Title: []string{"MH"}, Location: []string{"古代樹の森"}, Ranking: []monsters.Ranking{{Ranking: "1", VoteYear: "2024/3/12"}}},
		{Id: "0000000002", Name: "リオレイア", Description: "陸の女王", Category: "飛竜種", Title: []string{"MH"}, Location: []string{"古代樹の森"}, Ranking: []monsters.Ranking{{Ranking: "2", VoteYear: "2024/3/12"}}},
		{Id: "0000000003", Name: "ティガレックス", Description: "絶対強者", Category: "飛竜種", Title: []string{"MH"}, Location: []string{"古代樹の森"}, Ranking: []monsters.Ranking{{Ranking: "3", VoteYear: "2024/3/12"}}},
	}
	mysql.BeforeTest()
	conn := mysql.New(context.Background())
	type fields struct {
		qs *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*monsters.FetchMonsterRankingDto
		wantErr bool
	}{
		{name: "TestMonsterQueryService_FetchRank", fields: fields{qs: conn}, args: args{ctx: context.Background()}, want: wants, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MonsterQueryService{
				qs: tt.fields.qs,
			}
			got, err := s.FetchRank(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("MonsterQueryService.FetchRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonsterQueryService.FetchRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
