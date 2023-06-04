package gateway

import (
	"mh-api/api/driver"

	"github.com/stretchr/testify/mock"
)

type MockMonsterDriver struct {
	mock.Mock
}

func (m *MockMonsterDriver) GetAll() ([]driver.Monster) {
	args := m.Called()
	return args.Get(0).([]driver.Monster)
}

func (m *MockMonsterDriver) GetById(id int) (driver.Monster) {
	args := m.Called(id)
	return args.Get(0).(driver.Monster)
}

func (m *MockMonsterDriver) Create(monsterJson driver.MonsterJson) error {
	args := m.Called(monsterJson)
	return args.Error(0)
}