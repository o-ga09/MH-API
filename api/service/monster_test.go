package service

import (
	"mh-api/api/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T) {
	monsterInterface := new(MockMonsterInterface)
	service := MonsterService{monsterInterface}
	monsterInterface.On("GetAll").Return(entity.Monsters{},nil)
	actual,_ := service.GetAll()
	expected := entity.Monsters{}
	assert.Equal(t, expected, actual)
}

func Test_GetById(t *testing.T) {
	monsterInterface := new(MockMonsterInterface)
	service := MonsterService{monsterInterface}
	id := entity.MonsterId{Value: 1}
	monsterInterface.On("GetById",id).Return(entity.Monster{},nil)
	actual, _ := service.GetById(id)
	assert.Equal(t ,entity.Monster{},actual)
}

func Test_Create(t *testing.T) {
	monsterInterface := new(MockMonsterInterface)
	service := MonsterService{monsterInterface}
	monsterJson := entity.MonsterJson{Name: entity.MonsterName{Value: "ジンオウガ"},Desc: entity.MonsterDesc{Value: "渓流に生息する牙竜種"},Location: entity.MonsterLocation{Value: "渓流"},Specify: entity.MonsterSpecify{Value: "牙竜種"},Weakness_attack: entity.MonsterWeakness_A{Value: "頭部"},Weakness_element: entity.MonsterWeakness_E{Value: "氷"}}
	monsterInterface.On("Create",monsterJson).Return(nil)
	actual := service.Create(monsterJson)
	assert.Equal(t,nil,actual)
}

func Test_Update(t *testing.T) {
	mockInterface := new(MockMonsterInterface)
	service := MonsterService{mockInterface}
	id := entity.MonsterId{Value: 1}
	monsterJson := entity.MonsterJson{Name: entity.MonsterName{Value: "ジンオウガ"},Desc: entity.MonsterDesc{Value: "渓流に生息する牙竜種"},Location: entity.MonsterLocation{Value: "渓流"},Specify: entity.MonsterSpecify{Value: "牙竜種"},Weakness_attack: entity.MonsterWeakness_A{Value: "頭部"},Weakness_element: entity.MonsterWeakness_E{Value: "氷"}}
	mockInterface.On("Update",id,monsterJson).Return(nil)
	err := service.Update(id,monsterJson)
	assert.Equal(t,nil,err)
}