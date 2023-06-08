package service

import (
	"mh-api/api/entity"

	"github.com/stretchr/testify/mock"
)

type MockMonsterInterface struct {
	mock.Mock
}

func (m *MockMonsterInterface) GetAll() (entity.Monsters, error) {
	args := m.Called()
	return args.Get(0).(entity.Monsters), args.Error(1)
}

func (m *MockMonsterInterface) GetById(id entity.MonsterId) (entity.Monster, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Monster), args.Error(1)
}

func (m *MockMonsterInterface) Create(monsterJson entity.MonsterJson) error {
	args := m.Called(monsterJson)
	return args.Error(0)
}

func (m *MockMonsterInterface) Update(id entity.MonsterId,monsterJson entity.MonsterJson) error {
	args := m.Called(id,monsterJson)
	return args.Error(0)
}

func (m *MockMonsterInterface) Delete(id entity.MonsterId) error {
	args := m.Called(id)
	return args.Error(0)
}