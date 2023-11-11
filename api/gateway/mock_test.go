package gateway

import (
	"mh-api/api/gateway/repository"

	"github.com/stretchr/testify/mock"
)

type MockMonsterDriver struct {
	mock.Mock
}

func (m *MockMonsterDriver) GetAll() []repository.Monster {
	args := m.Called()
	return args.Get(0).([]repository.Monster)
}

func (m *MockMonsterDriver) GetById(id string) repository.Monster {
	args := m.Called(id)
	return args.Get(0).(repository.Monster)
}

func (m *MockMonsterDriver) Create(monsterJson repository.MonsterJson) error {
	args := m.Called(monsterJson)
	return args.Error(0)
}

func (m *MockMonsterDriver) Update(id string, monsterJson repository.MonsterJson) error {
	args := m.Called(id, monsterJson)
	return args.Error(0)
}

func (m *MockMonsterDriver) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
