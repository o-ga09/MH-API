package repository

import (
	"context"
	monsterdomain "mh-api/app/internal/domain/monsterDomain"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewMonsterRepository(t *testing.T) {
	mysql.BeforeTest()
	db := mysql.New(context.Background())
	type args struct {
		repo *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *MonsterRepository
	}{
		{name: "TestNewMonsterRepository", args: args{repo: db}, want: &MonsterRepository{repo: db}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonsterRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	db := mysql.New(context.Background())
	type fields struct {
		repo *gorm.DB
	}
	type args struct {
		ctx context.Context
		m   *monsterdomain.Monster
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestMonsterRepository_Save", fields: fields{repo: db}, args: args{ctx: context.Background(), m: monsterdomain.NewMonster("0000000004", "マムタロト", "周回対象")}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MonsterRepository{
				repo: tt.fields.repo,
			}
			if err := r.Save(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("MonsterRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMonsterRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	db := mysql.New(context.Background())
	type fields struct {
		repo *gorm.DB
	}
	type args struct {
		ctx       context.Context
		monsterId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestMonsterRepository_Remove", fields: fields{repo: db}, args: args{ctx: context.Background(), monsterId: "0000000004"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MonsterRepository{
				repo: tt.fields.repo,
			}
			if err := r.Remove(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("MonsterRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
