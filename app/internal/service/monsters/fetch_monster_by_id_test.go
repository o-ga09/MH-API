package monsters

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestNewFetchMonsterByID(t *testing.T) {
	qs := MonsterQueryServiceMock{}
	type args struct {
		qs MonsterQueryService
	}
	tests := []struct {
		name string
		args args
		want *FetchMonsterByIDService
	}{
		{name: "TestNewFetchMonsterByID", args: args{qs: &qs}, want: NewFetchMonsterByID(&qs)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFetchMonsterByID(tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFetchMonsterByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchMonsterByIDService_Run(t *testing.T) {
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
		want      *FetchMonsterListDto
		mockValue mockValue
		wantErr   bool
	}{
		{name: "TestFetchMonsterByIDService_Run", fields: fields{qs: &qs}, args: args{ctx: context.Background(), id: "1"}, want: &FetchMonsterListDto{}, mockValue: mockValue{res: nil, err: errors.New("Err")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs.FetchListFunc = func(ctx context.Context, id string) ([]*FetchMonsterListDto, error) {
				return tt.mockValue.res, tt.mockValue.err
			}
			f := FetchMonsterByIDService{
				qs: tt.fields.qs,
			}
			got, err := f.Run(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMonsterByIDService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchMonsterByIDService.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
