package service

import (
	"errors"
	"mh-api/api/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

var weakness = map[string]string{
	"頭部": "10",
	"前脚": "10",
	"胴体": "10",
	"後脚": "10",
	"尻尾": "10",
}

var response_1 = entity.Monster{
	Id:               entity.MonsterId{Value: 1},
	Name:             entity.MonsterName{Value: "ジンオウガ"},
	Desc:             entity.MonsterDesc{Value: "霊峰に生息する牙竜種"},
	Location:         entity.MonsterLocation{Value: "渓流"},
	Specify:          entity.MonsterSpecify{Value: "牙竜種"},
	Weakness_attack:  entity.MonsterWeakness_A{Value: weakness},
	Weakness_element: entity.MonsterWeakness_E{Value: weakness},
}

var response_2 = entity.Monster{
	Id:               entity.MonsterId{Value: 1},
	Name:             entity.MonsterName{Value: "ジンオウガ"},
	Desc:             entity.MonsterDesc{Value: "霊峰に生息する牙竜種"},
	Location:         entity.MonsterLocation{Value: "渓流"},
	Specify:          entity.MonsterSpecify{Value: "牙竜種"},
	Weakness_attack:  entity.MonsterWeakness_A{Value: weakness},
	Weakness_element: entity.MonsterWeakness_E{Value: weakness},
}
var monsterJson = entity.MonsterJson{
	Name:             entity.MonsterName{Value: "ジンオウガ"},
	Desc:             entity.MonsterDesc{Value: "渓流に生息する牙竜種"},
	Location:         entity.MonsterLocation{Value: "渓流"},
	Specify:          entity.MonsterSpecify{Value: "牙竜種"},
	Weakness_attack:  entity.MonsterWeakness_A{Value: weakness},
	Weakness_element: entity.MonsterWeakness_E{Value: weakness},
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
		{name: "正常系 - モンスターの詳細情報を取得可能な場合", args: entity.MonsterId{Value: 1}, expected: response_1, err: nil},
		{name: "異常系 - モンスターの詳細情報が取得不可な場合", args: entity.MonsterId{Value: 2}, expected: entity.Monster{}, err: errors.New("NOT FOUND")},
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
		{name: "正常系 - リストの要素が2つ以上の場合", args: monsterJson, err: nil},
		{name: "異常系 - リストの要素が空の場合", args: entity.MonsterJson{}, err: errors.New("FAILD TO REGISTERED")},
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
		{name: "正常系 - リストの要素が2つ以上の場合", arg1: entity.MonsterId{Value: 1}, arg2: monsterJson, err: nil},
		{name: "異常系 - リストの要素が空の場合", arg1: entity.MonsterId{Value: 2}, arg2: monsterJson, err: errors.New("FAILD TO UPDATED")},
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
		{name: "正常系 - リストの要素が2つ以上の場合", arg: entity.MonsterId{Value: 1}, err: nil},
		{name: "異常系 - リストの要素が空の場合", arg: entity.MonsterId{Value: 2}, err: errors.New("FAILD TO DELETED")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			monsterInterface := new(MockMonsterInterface)
			service := MonsterService{monsterInterface}
			monsterInterface.On("Delete",tt.arg).Return(tt.err)
			err := service.Delete(tt.arg)

			assert.Equal(t, tt.err, err)
		})
	}
}
