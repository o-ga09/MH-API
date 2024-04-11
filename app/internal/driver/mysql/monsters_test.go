package mysql

import (
	"context"
	"mh-api/app/internal/domain/monsters"
	"os"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewMonsterRepository(t *testing.T) {
	os.Setenv("DATABASE_URL", "mh-api:P@ssw0rd@tcp(127.0.0.1:3306)/mh-api?charset=utf8&parseTime=True&loc=Local")
	conn := New(context.Background())

	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *monsterRepository
	}{
		{name: "MonsetrRepository構造体を生成する", args: args{conn: conn}, want: NewMonsterRepository(conn)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monsterRepository_Get(t *testing.T) {
	BeforeTest()
	t.Cleanup(AfetrTest())
	conn := New(context.Background())

	wantMonster1 := monsters.NewMonster("0000000001", "リオレウス", "空の王者。")
	wantMonster2 := monsters.NewMonster("0000000002", "リオレイア", "陸の女王")
	wantMonster3 := monsters.NewMonster("0000000003", "ティガレックス", "絶対強者")

	wantMonsters1 := monsters.Monsters{
		wantMonster1, wantMonster2, wantMonster3,
	}

	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		monsterId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    monsters.Monsters
		wantErr bool
	}{
		{name: "DBからモンスターデータを複数件取得できる", fields: fields{conn: conn}, args: args{monsterId: ""}, want: wantMonsters1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &monsterRepository{
				conn: tt.fields.conn,
			}
			ctx := context.Background()
			got, err := r.Get(ctx, tt.args.monsterId)
			if (err != nil) != tt.wantErr {
				t.Errorf("monsterRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("monsterRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monsterRepository_Save(t *testing.T) {
	BeforeTest()
	t.Cleanup(AfetrTest())
	conn := New(context.Background())

	saveMonster1 := monsters.NewMonster("0000000004", "ライゼクス", "雷の反逆者")

	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		m monsters.Monster
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "DBからモンスターデータを複数件取得できる", fields: fields{conn: conn}, args: args{m: saveMonster1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &monsterRepository{
				conn: tt.fields.conn,
			}
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
	conn := New(context.Background())

	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		monsterId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "DBからモンスターデータを複数件取得できる", fields: fields{conn: conn}, args: args{monsterId: "0000000001"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &monsterRepository{
				conn: tt.fields.conn,
			}
			ctx := context.Background()
			if err := r.Remove(ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("monsterRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
