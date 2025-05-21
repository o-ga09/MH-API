package weakness

import (
	"context"
	"mh-api/app/internal/domain/weakness"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"
)

func TestNewweakRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	tests := []struct {
		name string

		want *weakRepository
	}{
		{name: "TestNewweakRepository", want: &weakRepository{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewweakRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewweakRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_weakRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	type fields struct {
	}
	type args struct {
		ctx context.Context
		w   weakness.Weakness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Save weakness successfully", args: args{ctx: context.Background(), w: *weakness.NewWeakness("0000000001", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test")}, wantErr: false},
		// Test case 2
		{name: "Save weakness with error", args: args{ctx: context.Background(), w: *weakness.NewWeakness("@$%&^#%$&&*%*&)(*)()", "", "", "", "", "", "", "", "", "", "", "", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &weakRepository{}
			if err := r.Save(tt.args.ctx, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("weakRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_weakRepository_Remove(t *testing.T) {
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
		{name: "Remove weakness successfully", args: args{ctx: context.Background(), monsterId: "1"}, wantErr: false},
		// Test case 2
		{name: "Remove weakness with error", args: args{ctx: context.Background(), monsterId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &weakRepository{}
			if err := r.Remove(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("weakRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
