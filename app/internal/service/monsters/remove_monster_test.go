package monsters

import (
	"context"
	monsterdomain "mh-api/app/internal/domain/monsterDomain"
	"reflect"
	"testing"
)

func TestNewRemoveMonsterService(t *testing.T) {
	repo := monsterdomain.MonsterRepositoryMock{}
	type args struct {
		repo monsterdomain.MonsterRepository
	}
	tests := []struct {
		name string
		args args
		want *RemoveMonsterService
	}{
		{name: "TestNewRemoveMonsterService", args: args{repo: &repo}, want: NewRemoveMonsterService(&repo)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRemoveMonsterService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRemoveMonsterService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveMonsterService_Run(t *testing.T) {
	repo := monsterdomain.MonsterRepositoryMock{}
	type fields struct {
		repo monsterdomain.MonsterRepository
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
		{name: "TestRemoveMonsterService_Run", fields: fields{repo: &repo}, args: args{ctx: context.Background(), monsterId: "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RemoveMonsterService{
				repo: tt.fields.repo,
			}
			repo.RemoveFunc = func(ctx context.Context, monsterId string) error {
				return nil
			}
			if err := r.Run(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("RemoveMonsterService.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
