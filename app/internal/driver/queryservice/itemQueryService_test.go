package queryservice

import (
	"context"
	"mh-api/app/internal/driver/mysql"
	"mh-api/app/internal/service/items"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewItemQueryService(t *testing.T) {
	mysql.BeforeTest()
	conn := mysql.New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *ItemQueryService
	}{
		{name: "TestNewItemQueryService", args: args{conn: conn}, want: NewItemQueryService(conn)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItemQueryService(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItemQueryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemQueryService_GetItems(t *testing.T) {
	mysql.BeforeTest()
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*items.ItemDto
		wantErr bool
	}{
		{name: "TestItemQueryService_GetItems", fields: fields{conn: conn}, want: []*items.ItemDto{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ItemQueryService{
				conn: tt.fields.conn,
			}
			got, err := s.GetItems(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemQueryService.GetItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == false && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemQueryService.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemQueryService_GetItem(t *testing.T) {
	mysql.BeforeTest()
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		itemId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *items.ItemDto
		wantErr bool
	}{
		{name: "TestItemQueryService_GetItem", fields: fields{conn: conn}, args: args{itemId: "1"}, want: &items.ItemDto{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ItemQueryService{
				conn: tt.fields.conn,
			}
			got, err := s.GetItem(context.Background(), tt.args.itemId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemQueryService.GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == false && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemQueryService.GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemQueryService_GetItemsByMonster(t *testing.T) {
	mysql.BeforeTest()
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    string
		want    *items.ItemsByMonster
		wantErr bool
	}{
		{name: "TestItemQueryService_GetItemsByMonster", fields: fields{conn: conn}, args: "1", want: &items.ItemsByMonster{}, wantErr: true},
		{name: "TestItemQueryService_GetItemsByMonster", fields: fields{conn: conn}, args: "0000000001", want: &items.ItemsByMonster{ItemId: "0000000001", ItemName: "リオレウスの鱗", Monster: []items.Monster{{ID: "0000000001", Name: "リオレウス"}}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ItemQueryService{
				conn: tt.fields.conn,
			}
			got, err := s.GetItemsByMonster(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemQueryService.GetItemsByMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == false && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemQueryService.GetItemsByMonster() = %v, want %v", got, tt.want)
			}
		})
	}
}
