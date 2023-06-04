package gateway

import (
	"mh-api/api/driver"
	"mh-api/api/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T) {
	mockDriver := new(MockMonsterDriver)
	gateway := MonsterGateway{mockDriver}
	monster := []driver.Monster{
		{Id: 1,Name: "ジンオウガ",Desc: "渓流に生息する牙竜種",Location: "渓流",Specify: "牙竜種",Weakness_attack: "頭部",Weakness_element: "氷"},
	}

	expected := entity.Monsters{
		Values: []entity.Monster{
			{Id: entity.MonsterId{Value: 1},Name: entity.MonsterName{Value: "ジンオウガ"},Desc: entity.MonsterDesc{Value: "渓流に生息する牙竜種"},Location: entity.MonsterLocation{Value: "渓流"},Specify: entity.MonsterSpecify{Value: "牙竜種"},Weakness_attack: entity.MonsterWeakness_A{Value: "頭部"},Weakness_element: entity.MonsterWeakness_E{Value: "氷"}},
		},
	}
	mockDriver.On("GetAll").Return(monster)
	actual, _ := gateway.GetAll()
	assert.Equal(t, expected, actual,)
}

func Test_GetById(t *testing.T)  {
	mockDriver := new(MockMonsterDriver)
	gateway := MonsterGateway{mockDriver}
	id := entity.MonsterId{Value: 1}
	monster := driver.Monster{Id: 1,Name: "ジンオウガ",Desc: "渓流に生息する牙竜種",Location: "渓流",Specify: "牙竜種",Weakness_attack: "頭部",Weakness_element: "氷"}
	mockDriver.On("GetById",1).Return(monster,nil)
	expected := entity.Monster{Id: entity.MonsterId{Value: 1},Name: entity.MonsterName{Value: "ジンオウガ"},Desc: entity.MonsterDesc{Value: "渓流に生息する牙竜種"},Location: entity.MonsterLocation{Value: "渓流"},Specify: entity.MonsterSpecify{Value: "牙竜種"},Weakness_attack: entity.MonsterWeakness_A{Value: "頭部"},Weakness_element: entity.MonsterWeakness_E{Value: "氷"}}
	actual, _ := gateway.GetById(id)
	assert.Equal(t,expected,actual)
}

func Test_Create(t *testing.T) {
	mockDriver := new(MockMonsterDriver)
	gateway := MonsterGateway{mockDriver}
	monsterJson := entity.MonsterJson{Name: entity.MonsterName{Value: "ジンオウガ"},Desc: entity.MonsterDesc{Value: "渓流に生息する牙竜種"},Location: entity.MonsterLocation{Value: "渓流"},Specify: entity.MonsterSpecify{Value: "牙竜種"},Weakness_attack: entity.MonsterWeakness_A{Value: "頭部"},Weakness_element: entity.MonsterWeakness_E{Value: "氷"}}
	driverJson := driver.MonsterJson{Name: "ジンオウガ",Desc: "渓流に生息する牙竜種",Location: "渓流",Specify: "牙竜種",Weakness_attack: "頭部",Weakness_element: "氷"}
	mockDriver.On("Create",driverJson).Return(nil)
	err := gateway.Create(monsterJson)
	assert.Equal(t,nil,err)
}

func Test_Update(t *testing.T) {
	mockDriver := new(MockMonsterDriver)
	gateway := MonsterGateway{mockDriver}
	monsterJson := entity.MonsterJson{Name: entity.MonsterName{Value: "ジンオウガ"},Desc: entity.MonsterDesc{Value: "渓流に生息する牙竜種"},Location: entity.MonsterLocation{Value: "渓流"},Specify: entity.MonsterSpecify{Value: "牙竜種"},Weakness_attack: entity.MonsterWeakness_A{Value: "頭部"},Weakness_element: entity.MonsterWeakness_E{Value: "氷"}}
	driverJson := driver.MonsterJson{Name: "ジンオウガ",Desc: "渓流に生息する牙竜種",Location: "渓流",Specify: "牙竜種",Weakness_attack: "頭部",Weakness_element: "氷"}
	id := entity.MonsterId{Value: 1}
	mockDriver.On("Update",1,driverJson).Return(nil)
	err := gateway.Update(id,monsterJson)
	assert.Equal(t,nil,err)
}