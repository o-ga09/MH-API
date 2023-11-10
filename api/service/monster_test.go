package service

import (
	"errors"
	"mh-api/api/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

var weakness_a = entity.Weakness_attack{
	FrontLegs: entity.AttackCatetgory{Slashing: "10", Blow: "10", Bullet: "10"},
	HindLegs:  entity.AttackCatetgory{Slashing: "10", Blow: "10", Bullet: "10"},
	Head:      entity.AttackCatetgory{Slashing: "10", Blow: "10", Bullet: "10"},
	Body:      entity.AttackCatetgory{Slashing: "10", Blow: "10", Bullet: "10"},
	Tail:      entity.AttackCatetgory{Slashing: "10", Blow: "10", Bullet: "10"},
}

var weakness_e = entity.Weakness_element{
	FrontLegs: entity.Elements{Fire: "10", Water: "10", Lightning: "10", Ice: "10", Dragon: "10"},
	HindLegs:  entity.Elements{Fire: "10", Water: "10", Lightning: "10", Ice: "10", Dragon: "10"},
	Head:      entity.Elements{Fire: "10", Water: "10", Lightning: "10", Ice: "10", Dragon: "10"},
	Body:      entity.Elements{Fire: "10", Water: "10", Lightning: "10", Ice: "10", Dragon: "10"},
	Tail:      entity.Elements{Fire: "10", Water: "10", Lightning: "10", Ice: "10", Dragon: "10"},
}

var response_1 = entity.Monster{
	Id:               entity.MonsterId{Value: "001"},
	Name:             entity.MonsterName{Value: "ジンオウガ"},
	Desc:             entity.MonsterDesc{Value: "霊峰に生息する牙竜種"},
	Location:         entity.MonsterLocation{Value: "渓流"},
	Category:         entity.MonsterCategory{Value: "牙竜種"},
	Title:            entity.GameTitle{Value: "モンスターハンターライズ"},
	Weakness_attack:  entity.MonsterWeakness_A{Value: weakness_a},
	Weakness_element: entity.MonsterWeakness_E{Value: weakness_e},
}

var response_2 = entity.Monster{
	Id:               entity.MonsterId{Value: "002"},
	Name:             entity.MonsterName{Value: "ジンオウガ"},
	Desc:             entity.MonsterDesc{Value: "霊峰に生息する牙竜種"},
	Location:         entity.MonsterLocation{Value: "渓流"},
	Category:         entity.MonsterCategory{Value: "牙竜種"},
	Title:            entity.GameTitle{Value: "モンスターハンターライズ"},
	Weakness_attack:  entity.MonsterWeakness_A{Value: weakness_a},
	Weakness_element: entity.MonsterWeakness_E{Value: weakness_e},
}
var monsterJson = entity.MonsterJson{
	Name:             entity.MonsterName{Value: "ジンオウガ"},
	Desc:             entity.MonsterDesc{Value: "渓流に生息する牙竜種"},
	Location:         entity.MonsterLocation{Value: "渓流"},
	Category:         entity.MonsterCategory{Value: "牙竜種"},
	Title:            entity.GameTitle{Value: "モンスターハンターライズ"},
	Weakness_attack:  entity.MonsterWeakness_A{Value: weakness_a},
	Weakness_element: entity.MonsterWeakness_E{Value: weakness_e},
}

func Test_GetAll(t *testing.T) {
	cases := []struct {
		name     string
		expected interface{}
		err      error
	}{
		{name: "正常系 - リストの要素が2つ以上の場合", expected: entity.Monsters{Values: []entity.Monster{response_1, response_2}}, err: nil},
		{name: "正常系 - リストの要素が1つの場合", expected: entity.Monsters{Values: []entity.Monster{response_1}}, err: nil},
		{name: "異常系 - リストの要素が空の場合", expected: entity.Monsters{}, err: errors.New("NOT FOUND")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			monsterInterface := new(MockMonsterInterface)
			service := MonsterService{monsterInterface}
			monsterInterface.On("GetAll").Return(tt.expected, tt.err)
			actual, err := service.GetAll()

			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_GetById(t *testing.T) {
	cases := []struct {
		name     string
		args     entity.MonsterId
		expected interface{}
		err      error
	}{
		{name: "正常系 - モンスターの詳細情報を取得可能な場合", args: entity.MonsterId{Value: "001"}, expected: response_1, err: nil},
		{name: "異常系 - モンスターの詳細情報が取得不可な場合", args: entity.MonsterId{Value: "002"}, expected: entity.Monster{}, err: errors.New("NOT FOUND")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			monsterInterface := new(MockMonsterInterface)
			service := MonsterService{monsterInterface}
			monsterInterface.On("GetById", tt.args).Return(tt.expected, tt.err)
			actual, err := service.GetById(tt.args)

			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_Create(t *testing.T) {
	cases := []struct {
		name string
		args entity.MonsterJson
		err  error
	}{
		{name: "正常系 - モンスターの詳細情報を登録成功", args: monsterJson, err: nil},
		{name: "異常系 - モンスターの詳細情報を登録失敗", args: entity.MonsterJson{}, err: errors.New("FAILD TO REGISTERED")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			monsterInterface := new(MockMonsterInterface)
			service := MonsterService{monsterInterface}
			monsterInterface.On("Create", monsterJson).Return(tt.err)
			err := service.Create(monsterJson)

			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_Update(t *testing.T) {
	cases := []struct {
		name string
		arg1 entity.MonsterId
		arg2 entity.MonsterJson
		err  error
	}{
		{name: "正常系 - モンスターの詳細情報を更新成功", arg1: entity.MonsterId{Value: "001"}, arg2: monsterJson, err: nil},
		{name: "異常系 - モンスターの詳細情報を更新失敗", arg1: entity.MonsterId{Value: "002"}, arg2: monsterJson, err: errors.New("FAILD TO UPDATED")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			monsterInterface := new(MockMonsterInterface)
			service := MonsterService{monsterInterface}
			monsterInterface.On("Update", tt.arg1, tt.arg2).Return(tt.err)
			err := service.Update(tt.arg1, tt.arg2)

			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_Delete(t *testing.T) {
	cases := []struct {
		name string
		arg  entity.MonsterId
		err  error
	}{
		{name: "正常系 - モンスターの詳細情報を削除成功", arg: entity.MonsterId{Value: "001"}, err: nil},
		{name: "異常系 - モンスターの詳細情報を削除失敗", arg: entity.MonsterId{Value: "002"}, err: errors.New("FAILD TO DELETED")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			monsterInterface := new(MockMonsterInterface)
			service := MonsterService{monsterInterface}
			monsterInterface.On("Delete", tt.arg).Return(tt.err)
			err := service.Delete(tt.arg)

			assert.Equal(t, tt.err, err)
		})
	}
}
