package items

import (
	"context"
	itemdomain "mh-api/app/internal/domain/itemDomain"
	"reflect"
	"testing"
)

func TestNewSaveItem(t *testing.T) {
	repo := itemdomain.ItemRepositoryMock{}
	type args struct {
		repo itemdomain.ItemRepository
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "TestNewSaveItem", args: args{repo: &repo}, want: NewSaveItem(&repo)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSaveItem(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSaveItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveItem_Run(t *testing.T) {
	repo := itemdomain.ItemRepositoryMock{}
	type fields struct {
		repo itemdomain.ItemRepository
	}
	type args struct {
		ctx  context.Context
		item *ItemDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestSaveItem_Run", fields: fields{repo: &repo}, args: args{ctx: context.Background(), item: &ItemDto{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.SaveFunc = func(ctx context.Context, m itemdomain.Item) error {
				return nil
			}
			s := &SaveItem{
				repo: tt.fields.repo,
			}
			if err := s.Run(tt.args.ctx, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("SaveItem.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
