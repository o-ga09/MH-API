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