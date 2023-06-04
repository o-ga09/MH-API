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