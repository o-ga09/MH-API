package tribes

import (
	"context"
	Tribes "mh-api/app/internal/domain/tribes"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewtribeRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *tribeRepository
	}{
		{name: "TestNewtribeRepository", args: args{conn: conn}, want: &tribeRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewtribeRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewtribeRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tribeRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		t   Tribes.Tribe
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Save tribe successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), t: *Tribes.NewTribe("0000000001", "test", "test", "test", "test")}, wantErr: false},
		// Test case 2
		{name: "Save tribe with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), t: *Tribes.NewTribe("@$%&^#%$&&*%*&)(*)()", "", "", "", "test")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &tribeRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("tribeRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tribeRepository_Remove(t *testing.T) {
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
		{name: "Remove tribe successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), monsterId: "1"}, wantErr: false},
		// Test case 2
		{name: "Remove tribe with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), monsterId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &tribeRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("tribeRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
