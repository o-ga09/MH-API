package items

import (
	"context"
	"mh-api/app/internal/domain/items"
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
		want *itemRepository
	}{
		{name: "TestNewMonsterRepository", args: args{conn: conn}, want: &itemRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		m   items.Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// test 1
		{name: "Save item successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), m: *items.NewItem("0000000001", "test", "test")}, wantErr: false},
		// test 2
		{name: "Save item with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), m: *items.NewItem("@$%&^#%$&&*%*&)(*)()", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &itemRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("itemRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_itemRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx    context.Context
		itemId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// test 1
		{name: "Remove item successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), itemId: "0000000001"}, wantErr: false},
		// test 2
		{name: "Remove item with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), itemId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &itemRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.itemId); (err != nil) != tt.wantErr {
				t.Errorf("itemRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
