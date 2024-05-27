package monsters

import (
	"context"
	monsterdomain "mh-api/app/internal/domain/monsterDomain"
	"reflect"
	"testing"
)

func TestNewSaveMonsterService(t *testing.T) {
	repo := monsterdomain.MonsterRepositoryMock{}
	type args struct {
		repo monsterdomain.MonsterRepository
	}
	tests := []struct {
		name string
		args args
		want *SaveMonsterService
	}{
		{name: "TestNewSaveMonsterService", args: args{repo: &repo}, want: NewSaveMonsterService(&repo)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSaveMonsterService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSaveMonsterService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveMonsterService_Run(t *testing.T) {
	repo := monsterdomain.MonsterRepositoryMock{}
	type fields struct {
		repo monsterdomain.MonsterRepository
	}
	type args struct {
		ctx context.Context
		m   *MonsterDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestSaveMonsterService_Run", fields: fields{repo: &repo}, args: args{ctx: context.Background(), m: &MonsterDto{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SaveMonsterService{
				repo: tt.fields.repo,
			}
			repo.SaveFunc = func(ctx context.Context, m *monsterdomain.Monster) error {
				return nil
			}
			if err := s.Run(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("SaveMonsterService.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
