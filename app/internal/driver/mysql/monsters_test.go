package mysql

import (
	"context"
	"mh-api/app/internal/domain/monsters"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewMonsterRepository(t *testing.T) {
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *monsterRepository
	}{
		// TODO: Add test cases.
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
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx       context.Context
		monsterId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    monsters.Monsters
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &monsterRepository{
				conn: tt.fields.conn,
			}
			got, err := r.Get(tt.args.ctx, tt.args.monsterId)
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
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		m   monsters.Monster
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &monsterRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("monsterRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_monsterRepository_Remove(t *testing.T) {
	type fields struct {
		conn *gorm.DB
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &monsterRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("monsterRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
