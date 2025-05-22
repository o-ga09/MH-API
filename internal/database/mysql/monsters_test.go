package mysql

import (
	"context"

	"mh-api/internal/domain/monsters"
	"reflect"
	"testing"
)

func TestNewMonsterRepository(t *testing.T) {
	BeforeTest()
	t.Cleanup(AfetrTest())

	tests := []struct {
		name string
		want *monsterRepository
	}{
		{name: "MonsetrRepository構造体を生成する", want: NewMonsterRepository()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monsterRepository_Save(t *testing.T) {
	BeforeTest()
	t.Cleanup(AfetrTest())

	saveMonster1 := monsters.NewMonster("0000000004", "ライゼクス", "雷の反逆者", "雷属性")

	type args struct {
		m monsters.Monster
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "DBからモンスターデータを複数件取得できる", args: args{m: saveMonster1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &monsterRepository{}
			ctx := context.Background()
			if err := r.Save(ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("monsterRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_monsterRepository_Remove(t *testing.T) {
	BeforeTest()
	t.Cleanup(AfetrTest())

	type args struct {
		monsterId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "DBからモンスターデータを削除できる", args: args{monsterId: "0000000001"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &monsterRepository{}
			ctx := context.Background()
			if err := r.Remove(ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("monsterRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
