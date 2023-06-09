package service

import (
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
	monsterJson := entity.MonsterJson{Name: entity.MonsterName{Value: "ジンオウガ"},Desc: entity.MonsterDesc{Value: "渓流に生息する牙竜種"},Location: entity.MonsterLocation{Value: "渓流"},Specify: entity.MonsterSpecify{Value: "牙竜種"},Weakness_attack: entity.MonsterWeakness_A{Value: weakness},Weakness_element: entity.MonsterWeakness_E{Value: weakness}}

	monsterInterface.On("Create",monsterJson).Return(nil)
	actual := service.Create(monsterJson)
	assert.Equal(t,nil,actual)
}

func Test_Update(t *testing.T) {
	mockInterface := new(MockMonsterInterface)
	service := MonsterService{mockInterface}
	id := entity.MonsterId{Value: 1}
	monsterJson := entity.MonsterJson{Name: entity.MonsterName{Value: "ジンオウガ"},Desc: entity.MonsterDesc{Value: "渓流に生息する牙竜種"},Location: entity.MonsterLocation{Value: "渓流"},Specify: entity.MonsterSpecify{Value: "牙竜種"},Weakness_attack: entity.MonsterWeakness_A{Value: weakness},Weakness_element: entity.MonsterWeakness_E{Value: weakness}}

	mockInterface.On("Update",id,monsterJson).Return(nil)
	err := service.Update(id,monsterJson)
	assert.Equal(t,nil,err)
}

func Test_Delete(t *testing.T) {
	mockInterface := new(MockMonsterInterface)
	service := MonsterService{mockInterface}
	id := entity.MonsterId{Value: 1}
	mockInterface.On("Delete",id).Return(nil)
	err := service.Delete(id)
	assert.Equal(t,nil,err)
}