package parts

import (
	"context"
	"mh-api/app/internal/domain/part"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"
)

func TestNewMonsterRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	tests := []struct {
		name string
		want *partRepository
	}{
		{name: "TestNewMonsterRepository", want: &partRepository{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	type args struct {
		ctx context.Context
		p   part.Part
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Save part successfully", args: args{ctx: context.Background(), p: *part.NewPart("0000000001", "test", "test")}, wantErr: false},
		// Test case 2
		{name: "Save part with error", args: args{ctx: context.Background(), p: *part.NewPart("@$%&^#%$&&*%*&)(*)()", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &partRepository{}
			if err := r.Save(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("partRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_partRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	type args struct {
		ctx       context.Context
		monsterId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Remove part successfully", args: args{ctx: context.Background(), monsterId: "0000000001"}, wantErr: false},
		// Test case 2
		{name: "Remove part with error", args: args{ctx: context.Background(), monsterId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &partRepository{}
			if err := r.Remove(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("partRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
