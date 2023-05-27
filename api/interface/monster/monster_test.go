package monster

import (
	"context"
	"mh-api/api/entity"
	"mh-api/api/interface/repository"
	"reflect"
	"testing"
)

func TestNewMosterService(t *testing.T) {
	t.Skip()
	type args struct {
		r repository.IMonsterRepository
	}
	tests := []struct {
		name string
		args args
		want IMonsterService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMosterService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMosterService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monster_FindAllMonsters(t *testing.T) {
	t.Skip()
	type fields struct {
		repo repository.IMonsterRepository
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
			s := &monster{
				repo: tt.fields.repo,
			}
			got, err := s.FindAllMonsters(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("monster.FindAllMonsters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("monster.FindAllMonsters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monster_FindMonsterById(t *testing.T) {
	t.Skip()
	type fields struct {
		repo repository.IMonsterRepository
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
			s := &monster{
				repo: tt.fields.repo,
			}
			got, err := s.FindMonsterById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("monster.FindMonsterById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("monster.FindMonsterById() = %v, want %v", got, tt.want)
			}
		})
	}
}
