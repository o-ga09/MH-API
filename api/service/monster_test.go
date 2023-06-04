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