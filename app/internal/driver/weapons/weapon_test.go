package weapons

import (
	"context"
	"mh-api/app/internal/domain/weapons"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewweaponRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type args struct {
		conn *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *weaponRepository
	}{
		{name: "TestNewweaponRepository", args: args{conn: conn}, want: &weaponRepository{conn: conn}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewweaponRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewweaponRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_weaponRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx context.Context
		w   weapons.Weapon
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Save weapon successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), w: *weapons.NewWeapon("weaponId", "name", "imageUrl", "rerarity", "attack", "elementAttac", "shapness", "critical", "description")}, wantErr: false},
		// Test case 2
		{name: "Save weapon with empty", fields: fields{conn: conn}, args: args{ctx: context.Background(), w: *weapons.NewWeapon("", "", "", "", "", "", "", "", "")}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &weaponRepository{
				conn: tt.fields.conn,
			}
			if err := r.Save(tt.args.ctx, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("weaponRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_weaponRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())
	conn := mysql.New(context.Background())
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		ctx      context.Context
		weaponId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Test case 1
		{name: "Remove weapon successfully", fields: fields{conn: conn}, args: args{ctx: context.Background(), weaponId: "1"}, wantErr: false},
		// Test case 2
		{name: "Remove weapon with error", fields: fields{conn: conn}, args: args{ctx: context.Background(), weaponId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &weaponRepository{
				conn: tt.fields.conn,
			}
			if err := r.Remove(tt.args.ctx, tt.args.weaponId); (err != nil) != tt.wantErr {
				t.Errorf("weaponRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
