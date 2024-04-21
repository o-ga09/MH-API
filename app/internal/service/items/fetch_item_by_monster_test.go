package items

import (
	"context"
	"reflect"
	"testing"
)

func TestNewFetchItemByMonster(t *testing.T) {
	qs := ItemQueryServiceMock{}
	type args struct {
		qs ItemQueryService
	}
	tests := []struct {
		name string
		args args
		want *FetchItemByMonster
	}{
		{name: "TestNewFetchItemByMonster", args: args{qs: &qs}, want: NewFetchItemByMonster(&qs)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFetchItemByMonster(tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFetchItemByMonster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchItemByMonster_Run(t *testing.T) {
	qs := ItemQueryServiceMock{}
	type fields struct {
		qs ItemQueryService
	}
	type args struct {
		ctx    context.Context
		itemId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ItemsByMonster
		wantErr bool
	}{
		{name: "TestFetchItemByMonster_Run", fields: fields{qs: &qs}, args: args{ctx: context.Background(), itemId: ""}, want: &ItemsByMonster{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs.GetItemsByMonsterFunc = func(ctx context.Context, itemId string) (*ItemsByMonster, error) {
				return &ItemsByMonster{}, nil
			}
			s := &FetchItemByMonster{
				qs: tt.fields.qs,
			}
			got, err := s.Run(tt.args.ctx, tt.args.itemId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchItemByMonster.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchItemByMonster.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
