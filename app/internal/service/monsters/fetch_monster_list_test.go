package monsters

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestNewFetchMonsterListService(t *testing.T) {
	qs := MonsterQueryServiceMock{}
	type args struct {
		qs MonsterQueryService
	}
	tests := []struct {
		name string
		args args
		want *FetchMonsterListService
	}{
		{name: "TestNewFetchMonsterListService", args: args{qs: &qs}, want: NewFetchMonsterListService(&qs)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFetchMonsterListService(tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFetchMonsterListService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchMonsterListService_Run(t *testing.T) {
	qs := MonsterQueryServiceMock{}
	type fields struct {
		qs MonsterQueryService
	}
	type args struct {
		ctx context.Context
		id  string
	}
	type mockValue struct {
		res []*FetchMonsterListDto
		err error
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mockValue mockValue
		want      []*FetchMonsterListDto
		wantErr   bool
	}{
		{name: "TestFetchMonsterListService_Run", fields: fields{qs: &qs}, args: args{ctx: context.Background(), id: "1"}, mockValue: mockValue{res: []*FetchMonsterListDto{}, err: errors.New("Err")}, want: []*FetchMonsterListDto{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FetchMonsterListService{
				qs: tt.fields.qs,
			}
			qs.FetchListFunc = func(ctx context.Context, id string) ([]*FetchMonsterListDto, error) {
				return tt.mockValue.res, tt.mockValue.err
			}
			got, err := s.Run(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMonsterListService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchMonsterListService.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
