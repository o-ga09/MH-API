package repository

import (
	"context"
	itemdomain "mh-api/app/internal/domain/itemDomain"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewItemRepository(t *testing.T) {
	mysql.BeforeTest()
	conn := mysql.New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *ItemRepository
	}{
		{name: "TestNewItemRepository", args: args{conn: conn}, want: &ItemRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItemRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		m   itemdomain.Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestItemRepository_Save", fields: fields{conn: conn}, args: args{ctx: context.Background(), m: itemdomain.Item{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ItemRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("ItemRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestItemRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
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
		{name: "TestItemRepository_Remove", fields: fields{conn: conn}, args: args{ctx: context.Background(), itemId: "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ItemRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.itemId); (err != nil) != tt.wantErr {
				t.Errorf("ItemRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
