package parts

import (
	"context"
	"mh-api/app/internal/domain/part"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewMonsterRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *partRepository
	}{
		{name: "TestNewMonsterRepository", args: args{conn: conn}, want: &partRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		p   part.Part
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Save part successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), p: *part.NewPart("0000000001", "test", "test", "test")}, wantErr: false},
		// Test case 2
		{name: "Save part with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), p: *part.NewPart("@$%&^#%$&&*%*&)(*)()", "", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &partRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("partRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_partRepository_Remove(t *testing.T) {
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
		{name: "Remove part successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), monsterId: "0000000001"}, wantErr: false},
		// Test case 2
		{name: "Remove part with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), monsterId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &partRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("partRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
