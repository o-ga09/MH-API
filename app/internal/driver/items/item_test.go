package items

import (
	"context"
	"mh-api/app/internal/domain/items"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"
)

func TestNewMonsterRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	tests := []struct {
		name string
		want *itemRepository
	}{
		{name: "TestNewMonsterRepository", want: &itemRepository{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonsterRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonsterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	type args struct {
		ctx context.Context
		m   items.Item
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// test 1
		{name: "Save item successfully", args: args{ctx: context.Background(), m: *items.NewItem("0000000001", "test", "test")}, wantErr: false},
		// test 2
		{name: "Save item with error", args: args{ctx: context.Background(), m: *items.NewItem("@$%&^#%$&&*%*&)(*)()", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &itemRepository{}
			if err := r.Save(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("itemRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_itemRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	type args struct {
		ctx    context.Context
		itemId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// test 1
		{name: "Remove item successfully", args: args{ctx: context.Background(), itemId: "0000000001"}, wantErr: false},
		// test 2
		{name: "Remove item with error", args: args{ctx: context.Background(), itemId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &itemRepository{}
			if err := r.Remove(tt.args.ctx, tt.args.itemId); (err != nil) != tt.wantErr {
				t.Errorf("itemRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
