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