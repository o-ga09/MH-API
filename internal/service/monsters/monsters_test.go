package monsters_test

import (
	"context"
	"fmt"
	monsterDomain "mh-api/internal/domain/monsters"
	monsterService "mh-api/internal/service/monsters"

	"reflect"
	"testing"
)

func TestNewMonsterService(t *testing.T) {
	repo := monsterDomain.RepositoryMock{}
	qs := monsterService.MonsterQueryServiceMock{}
	type args struct {
		repo monsterDomain.Repository
		qs   monsterService.MonsterQueryService
	}
	tests := []struct {
		name string
		args args
		want *monsterService.MonsterService
	}{
		{name: "test_new_monster_service", args: args{repo: &repo, qs: &qs}, want: monsterService.NewMonsterService(&repo, &qs)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monsterService.NewMonsterService(tt.args.repo, tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonsterService_SaveMonster(t *testing.T) {
	repo := monsterDomain.RepositoryMock{}
	qs := monsterService.MonsterQueryServiceMock{}

	SaveData := monsterService.MonsterDto{
		ID: "0000000001", Name: "リオレウス", Description: "空の全てを統べる王者。",
	}

	type fields struct {
		repo monsterDomain.Repository
		qs   monsterService.MonsterQueryService
	}

	type mockValue struct {
		err error
	}
	type args struct {
		param monsterService.MonsterDto
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mockValue mockValue
		wantErr   bool
	}{
		{name: "モンスターテーブルのデータを更新できる", fields: fields{repo: &repo, qs: &qs}, args: args{param: SaveData}, mockValue: mockValue{err: nil}, wantErr: false},
		{name: "モンスターテーブルのデータを更新できない", fields: fields{repo: &repo, qs: &qs}, args: args{param: SaveData}, mockValue: mockValue{err: fmt.Errorf("can not get record")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.SaveFunc = func(ctx context.Context, m monsterDomain.Monster) error {
				return tt.mockValue.err
			}

			ctx := context.Background()
			s := monsterService.NewMonsterService(tt.fields.repo, tt.fields.qs)
			err := s.SaveMonsters(ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("MonsterService.GetMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMonsterService_RemoveMonster(t *testing.T) {
	repo := monsterDomain.RepositoryMock{}
	qs := monsterService.MonsterQueryServiceMock{}

	type fields struct {
		repo monsterDomain.Repository
		qs   monsterService.MonsterQueryService
	}

	type mockValue struct {
		err error
	}
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mockValue mockValue
		want      []*monsterService.MonsterDto
		wantErr   bool
	}{
		{name: "モンスターテーブルのデータを削除できる", fields: fields{repo: &repo, qs: &qs}, args: args{id: "0000000001"}, mockValue: mockValue{err: nil}, wantErr: false},
		{name: "モンスターテーブルのデータを削除できない", fields: fields{repo: &repo, qs: &qs}, args: args{id: ""}, mockValue: mockValue{err: fmt.Errorf("can not get record")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.RemoveFunc = func(ctx context.Context, monsterId string) error {
				return tt.mockValue.err
			}
			ctx := context.Background()
			s := monsterService.NewMonsterService(tt.fields.repo, tt.fields.qs)
			err := s.RemoveMonsters(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MonsterService.GetMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMonsterService_FetchMonsterDetail(t *testing.T) {
	repo := monsterDomain.RepositoryMock{}
	qs := monsterService.MonsterQueryServiceMock{}

	weak_A := []monsterService.Weakness_attack{
		{PartId: "001", Slashing: "45", Blow: "45", Bullet: "45"},
	}
	weak_E := []monsterService.Weakness_element{
		{PartId: "001", Fire: "45", Water: "45", Thunder: "45", Ice: "45", Dragon: "45"},
	}
	wantMonsters := []*monsterService.FetchMonsterListDto{
		{Id: "0000000001", Name: "リオレウス", Description: "空の全てを統べる王者。", Location: []string{"渓流", "古代樹の森"}, Category: "飛竜種", Title: []string{"MH", "MH2"}, FirstWeak_Attack: "頭部", SecondWeak_Attack: "翼", FirstWeak_Element: "龍", SecondWeak_Element: "雷", Weakness_attack: weak_A, Weakness_element: weak_E},
	}
	type fields struct {
		repo monsterDomain.Repository
		qs   monsterService.MonsterQueryService
	}
	type args struct {
		ctx context.Context
		id  string
	}
	type mockValue struct {
		value []*monsterService.FetchMonsterListDto
		err   error
	}

	var expectedMockValues []*monsterService.FetchMonsterListDto
	v1 := monsterService.FetchMonsterListDto{
		Id:                 "0000000001",
		Name:               "リオレウス",
		Description:        "空の全てを統べる王者。",
		Location:           []string{"渓流", "古代樹の森"},
		Category:           "飛竜種",
		Title:              []string{"MH", "MH2"},
		FirstWeak_Attack:   "頭部",
		SecondWeak_Attack:  "翼",
		FirstWeak_Element:  "龍",
		SecondWeak_Element: "雷",
		Weakness_attack:    weak_A,
		Weakness_element:   weak_E,
	}
	expectedMockValues = append(expectedMockValues, &v1)

	tests := []struct {
		name      string
		fields    fields
		args      args
		mockValue mockValue
		want      []*monsterService.FetchMonsterListDto
		wantErr   bool
	}{
		{name: "モンスター検索結果を取得する(複数件)", fields: fields{repo: &repo, qs: &qs}, args: args{ctx: context.Background(), id: ""}, mockValue: mockValue{value: expectedMockValues, err: nil}, want: wantMonsters, wantErr: false},
		{name: "モンスター検索結果を取得する(1件)", fields: fields{repo: &repo, qs: &qs}, args: args{ctx: context.Background(), id: "0000000001"}, mockValue: mockValue{value: expectedMockValues, err: nil}, want: wantMonsters, wantErr: false},
		{name: "モンスター検索結果を取得できない", fields: fields{repo: &repo, qs: &qs}, args: args{ctx: context.Background(), id: ""}, mockValue: mockValue{value: nil, err: fmt.Errorf("can not get record")}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs.FetchListFunc = func(ctx context.Context, id string) ([]*monsterService.FetchMonsterListDto, error) {
				return tt.mockValue.value, tt.mockValue.err
			}
			s := monsterService.NewMonsterService(tt.fields.repo, tt.fields.qs)
			got, err := s.FetchMonsterDetail(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MonsterService.FetchMonsterDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonsterService.FetchMonsterDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
