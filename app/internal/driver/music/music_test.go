package music

import (
	"context"
	"mh-api/app/internal/domain/music"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewmusicRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *musicRepository
	}{
		{name: "TestNewmusicRepository", args: args{conn: conn}, want: &musicRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMusicRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewmusicRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_musicRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		m   music.Music
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Save music successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), m: *music.NewMusic("0000000001", "0000000001", "test", "test")}, wantErr: false},
		// Test case 2
		{name: "Save music with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), m: *music.NewMusic("@$%&^#%$&&*%*&)(*)()", "", "", "")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &musicRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("musicRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_musicRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx   context.Context
		bgmId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Remove music successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), bgmId: "0000000001"}, wantErr: false},
		// Test case 2
		{name: "Remove music with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), bgmId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &musicRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.bgmId); (err != nil) != tt.wantErr {
				t.Errorf("musicRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
