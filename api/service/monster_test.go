package service

import (
	"mh-api/api/entity"
	"mh-api/api/interface/monster"
	"reflect"
	"testing"

	"golang.org/x/net/context"
)

func TestNewMonsterUsecase(t *testing.T) {
	type args struct {
		u monster.IMonsterService
	}
	tests := []struct {
		name string
		args args
		want MonsterService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterUsecase(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monsterService_FindAllMonsters(t *testing.T) {
	t.Skip()
	type fields struct {
		use monster.IMonsterService
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Monsters
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &monsterService{
				use: tt.fields.use,
			}
			got, err := u.FindAllMonsters(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("monsterService.FindAllMonsters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("monsterService.FindAllMonsters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monsterService_FindMonsterById(t *testing.T) {
	type fields struct {
		use monster.IMonsterService
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Monster
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &monsterService{
				use: tt.fields.use,
			}
			got, err := u.FindMonsterById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("monsterService.FindMonsterById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("monsterService.FindMonsterById() = %v, want %v", got, tt.want)
			}
		})
	}
}
