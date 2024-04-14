package fields

import (
	"context"
	fieldsDomain "mh-api/app/internal/domain/fields"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewfieldRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())

	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *fieldRepository
	}{
		{name: "TestNewfieldRepository", args: args{conn: conn}, want: &fieldRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewfieldRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewfieldRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())

	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		f   fieldsDomain.Field
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Save field successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), f: *fieldsDomain.NewField("0000000001", "0000000001", "test", "test")}, wantErr: false},
		// Test case 2
		{name: "Save field with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), f: *fieldsDomain.NewField("", "", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &fieldRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("fieldRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fieldRepository_Remove(t *testing.T) {
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
		{name: "Remove field successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), monsterId: "0000000001"}, wantErr: false},
		// Test case 2
		{name: "Remove field with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), monsterId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &fieldRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.monsterId); (err != nil) != tt.wantErr {
				t.Errorf("fieldRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
