package ranking

import (
	"context"
	"mh-api/app/internal/domain/ranking"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"
)

func TestNewMonsterRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	tests := []struct {
		name string
		want *rankingRepository
	}{
		{name: "TestNewMonsterRepository", want: &rankingRepository{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rankingRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	type fields struct {
	}
	type args struct {
		ctx  context.Context
		rank ranking.Ranking
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Save ranking successfully", args: args{ctx: context.Background(), rank: *ranking.NewRanking("0000000001", "test", "test")}, wantErr: false},
		// Test case 2
		{name: "Save ranking with error", args: args{ctx: context.Background(), rank: *ranking.NewRanking("@$%&^#%$&&*%*&)(*)()", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rankingRepository{}
			if err := r.Save(tt.args.ctx, tt.args.rank); (err != nil) != tt.wantErr {
				t.Errorf("rankingRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_rankingRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	type fields struct {
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
		// Test case 1
		{name: "Remove ranking successfully", args: args{ctx: context.Background(), monsterId: "0000000001"}, wantErr: false},
		// Test case 2
		{name: "Remove ranking with error", args: args{ctx: context.Background(), monsterId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rankingRepository{}
			if err := r.Remove(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("rankingRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
