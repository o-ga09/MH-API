package weapons

import (
	"context"
	"mh-api/app/internal/domain/weapons"
	"mh-api/app/internal/driver/mysql"
	"reflect"
	"testing"
)

func TestNewweaponRepository(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	tests := []struct {
		name string

		want *weaponRepository
	}{
		{name: "TestNewweaponRepository", want: &weaponRepository{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewweaponRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewweaponRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_weaponRepository_Save(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	type fields struct {
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
		{name: "Save weapon successfully", args: args{ctx: context.Background(), w: *weapons.NewWeapon("weaponId", "name", "imageUrl", "rerarity", "attack", "elementAttac", "shapness", "critical", "description")}, wantErr: false},
		// Test case 2
		{name: "Save weapon with empty", args: args{ctx: context.Background(), w: *weapons.NewWeapon("", "", "", "", "", "", "", "", "")}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &weaponRepository{}
			if err := r.Save(tt.args.ctx, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("weaponRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_weaponRepository_Remove(t *testing.T) {
	mysql.BeforeTest()
	t.Cleanup(mysql.AfetrTest())

	type fields struct {
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
		{name: "Remove weapon successfully", args: args{ctx: context.Background(), weaponId: "1"}, wantErr: false},
		// Test case 2
		{name: "Remove weapon with error", args: args{ctx: context.Background(), weaponId: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &weaponRepository{}
			if err := r.Remove(tt.args.ctx, tt.args.weaponId); (err != nil) != tt.wantErr {
				t.Errorf("weaponRepository.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
