package items

import (
	"context"
	itemdomain "mh-api/app/internal/domain/itemDomain"
	"reflect"
	"testing"
)

func TestNewRemoveItem(t *testing.T) {
	repo := itemdomain.ItemRepositoryMock{}
	type args struct {
		repo itemdomain.ItemRepository
	}
	tests := []struct {
		name string
		args args
		want *RemoveItem
	}{
		{name: "TestNewRemoveItem", args: args{repo: &repo}, want: NewRemoveItem(&repo)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRemoveItem(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRemoveItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveItem_Run(t *testing.T) {
	repo := itemdomain.ItemRepositoryMock{}
	type fields struct {
		repo itemdomain.ItemRepository
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
		{name: "TestRemoveItem_Run", fields: fields{repo: &repo}, args: args{ctx: context.Background(), itemId: "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.RemoveFunc = func(ctx context.Context, id string) error {
				return nil
			}
			s := &RemoveItem{
				repo: tt.fields.repo,
			}
			if err := s.Run(tt.args.ctx, tt.args.itemId); (err != nil) != tt.wantErr {
				t.Errorf("RemoveItem.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
