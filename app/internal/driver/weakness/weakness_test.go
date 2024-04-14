package weakness

import (
	"context"
	"mh-api/app/internal/domain/weakness"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewweakRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *weakRepository
	}{
		{name: "TestNewweakRepository", args: args{conn: conn}, want: &weakRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewweakRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewweakRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_weakRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
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
		{name: "Save weakness successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), w: *weakness.NewWeakness("0000000001", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test")}, wantErr: false},
		// Test case 2
		{name: "Save weakness with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), w: *weakness.NewWeakness("@$%&^#%$&&*%*&)(*)()", "", "", "", "", "", "", "", "", "", "", "", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &weakRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("weakRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_weakRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
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
		// Test case 1
		{name: "Remove weakness successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), monsterId: "1"}, wantErr: false},
		// Test case 2
		{name: "Remove weakness with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), monsterId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &weakRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("weakRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
