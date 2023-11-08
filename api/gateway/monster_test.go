package gateway

import (
	"errors"
	"mh-api/api/driver"
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

var mock_response_1 = driver.Monster{
	Id: 1, 
	Name: "ジンオウガ", 
	Desc: "霊峰に生息する牙竜種", 
	Location: "渓流", 
	Specify: "牙竜種", 
	Weakness_attack: "10 10 10 10 10", 
	Weakness_element: "10 10 10 10 10",
}

var mock_response_2 = driver.Monster{
	Id: 2, 
	Name: "ジンオウガ", 
	Desc: "霊峰に生息する牙竜種", 
	Location: "渓流", 
	Specify: "牙竜種", 
	Weakness_attack: "10 10 10 10 10", 
	Weakness_element: "10 10 10 10 10",
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
	Id:               entity.MonsterId{Value: 2},
	Name:             entity.MonsterName{Value: "ジンオウガ"},
	Desc:             entity.MonsterDesc{Value: "霊峰に生息する牙竜種"},
	Location:         entity.MonsterLocation{Value: "渓流"},
	Specify:          entity.MonsterSpecify{Value: "牙竜種"},
	Weakness_attack:  entity.MonsterWeakness_A{Value: weakness},
	Weakness_element: entity.MonsterWeakness_E{Value: weakness},
}

var mock_arg = driver.MonsterJson{
	Name: "ジンオウガ", 
	Desc: "霊峰に生息する牙竜種", 
	Location: "渓流", 
	Specify: "牙竜種", 
	Weakness_attack: "10 10 10 10 10 ", 
	Weakness_element: "10 10 10 10 10 ",
}

var arg = entity.MonsterJson{
	Name:             entity.MonsterName{Value: "ジンオウガ"},
	Desc:             entity.MonsterDesc{Value: "霊峰に生息する牙竜種"},
	Location:         entity.MonsterLocation{Value: "渓流"},
	Specify:          entity.MonsterSpecify{Value: "牙竜種"},
	Weakness_attack:  entity.MonsterWeakness_A{Value: weakness},
	Weakness_element: entity.MonsterWeakness_E{Value: weakness},
}

func Test_GetAll(t *testing.T) {
	cases := []struct {
		name     string
		mock interface{}
		expected interface{}
		err      error
	}{
		{name: "正常系 - リストの要素が2つ以上の場合", mock: []driver.Monster{mock_response_1, mock_response_2},expected: entity.Monsters{Values: []entity.Monster{response_1,response_2}} ,err: nil},
		{name: "正常系 - リストの要素が1つの場合", mock: []driver.Monster{mock_response_1},expected: entity.Monsters{Values: []entity.Monster{response_1}} ,err: nil},
		{name: "異常系 - リストの要素が空の場合", mock: []driver.Monster{},expected: entity.Monsters{} ,err: errors.New("0件のレコードを取得しました！")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockDriver := new(MockMonsterDriver)
			gateway := MonsterGateway{mockDriver}
			mockDriver.On("GetAll").Return(tt.mock, tt.err)
			actual, err := gateway.GetAll()

			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_GetById(t *testing.T) {
	cases := []struct {
		name     string
		arg entity.MonsterId
		mockarg int
		mock interface{}
		expected interface{}
		err      error
	}{
		{name: "正常系 - モンスターの詳細情報を取得可能な場合", arg: entity.MonsterId{Value: 1},mockarg: 1,mock: mock_response_1,expected: response_1 ,err: nil},
		{name: "異常系 - モンスターの詳細情報が取得不可な場合", arg: entity.MonsterId{Value: 2},mockarg: 2,mock: driver.Monster{},expected: entity.Monster{} ,err: errors.New("id = {2} のレコードはありませんでした！")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockDriver := new(MockMonsterDriver)
			gateway := MonsterGateway{mockDriver}
			mockDriver.On("GetById",tt.mockarg).Return(tt.mock, tt.err)
			actual, err := gateway.GetById(tt.arg)

			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_Create(t *testing.T) {
	cases := []struct {
		name     string
		arg entity.MonsterJson
		mockarg interface{}
		err      error
	}{
		{name: "正常系 - モンスターの詳細情報を登録成功",arg: arg, mockarg: mock_arg ,err: nil},
		{name: "異常系 - モンスターの詳細情報を登録失敗",arg: entity.MonsterJson{}, mockarg: driver.MonsterJson{} ,err: errors.New("FAILD TO REGISTERED")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockDriver := new(MockMonsterDriver)
			gateway := MonsterGateway{mockDriver}
			mockDriver.On("Create",tt.mockarg).Return(tt.err)
			err := gateway.Create(tt.arg)

			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_Update(t *testing.T) {
	cases := []struct {
		name     string
		arg1 entity.MonsterId
		arg2 entity.MonsterJson
		mockarg1 int
		mockarg2 interface{}
		err      error
	}{
		{name: "正常系 - モンスターの詳細情報を更新成功",arg1: entity.MonsterId{Value: 1},arg2: arg,mockarg1: 1, mockarg2: mock_arg ,err: nil},
		{name: "異常系 - モンスターの詳細情報を更新失敗",arg1: entity.MonsterId{Value: 2},arg2: entity.MonsterJson{},mockarg1: 2, mockarg2: driver.MonsterJson{} ,err: errors.New("FAILED TO UPDATED")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockDriver := new(MockMonsterDriver)
			gateway := MonsterGateway{mockDriver}
			mockDriver.On("Update",tt.mockarg1,tt.mockarg2).Return(tt.err)
			err := gateway.Update(tt.arg1,tt.arg2)

			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_Delete(t *testing.T) {
	cases := []struct {
		name     string
		arg entity.MonsterId
		mockarg int
		err      error
	}{
		{name: "正常系 - モンスターの詳細情報を削除成功",arg: entity.MonsterId{Value: 1} ,mockarg: 1 ,err: nil},
		{name: "異常系 - モンスターの詳細情報を削除失敗", arg: entity.MonsterId{Value: 2},mockarg: 2,err: errors.New("FAILED TO DELETED")},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockDriver := new(MockMonsterDriver)
			gateway := MonsterGateway{mockDriver}
			mockDriver.On("Delete",tt.mockarg).Return( tt.err)
			err := gateway.Delete(tt.arg)

			assert.Equal(t, tt.err, err)
		})
	}
}
