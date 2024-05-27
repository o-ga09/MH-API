// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package monsterdomain

import (
	"context"
	"sync"
)

// Ensure, that MonsterRepositoryMock does implement MonsterRepository.
// If this is not the case, regenerate this file with moq.
var _ MonsterRepository = &MonsterRepositoryMock{}

// MonsterRepositoryMock is a mock implementation of MonsterRepository.
//
//	func TestSomethingThatUsesMonsterRepository(t *testing.T) {
//
//		// make and configure a mocked MonsterRepository
//		mockedMonsterRepository := &MonsterRepositoryMock{
//			RemoveFunc: func(ctx context.Context, monsterId string) error {
//				panic("mock out the Remove method")
//			},
//			SaveFunc: func(ctx context.Context, m *Monster) error {
//				panic("mock out the Save method")
//			},
//		}
//
//		// use mockedMonsterRepository in code that requires MonsterRepository
//		// and then make assertions.
//
//	}
type MonsterRepositoryMock struct {
	// RemoveFunc mocks the Remove method.
	RemoveFunc func(ctx context.Context, monsterId string) error

	// SaveFunc mocks the Save method.
	SaveFunc func(ctx context.Context, m *Monster) error

	// calls tracks calls to the methods.
	calls struct {
		// Remove holds details about calls to the Remove method.
		Remove []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// MonsterId is the monsterId argument value.
			MonsterId string
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// M is the m argument value.
			M *Monster
		}
	}
	lockRemove sync.RWMutex
	lockSave   sync.RWMutex
}

// Remove calls RemoveFunc.
func (mock *MonsterRepositoryMock) Remove(ctx context.Context, monsterId string) error {
	if mock.RemoveFunc == nil {
		panic("MonsterRepositoryMock.RemoveFunc: method is nil but MonsterRepository.Remove was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		MonsterId string
	}{
		Ctx:       ctx,
		MonsterId: monsterId,
	}
	mock.lockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	mock.lockRemove.Unlock()
	return mock.RemoveFunc(ctx, monsterId)
}

// RemoveCalls gets all the calls that were made to Remove.
// Check the length with:
//
//	len(mockedMonsterRepository.RemoveCalls())
func (mock *MonsterRepositoryMock) RemoveCalls() []struct {
	Ctx       context.Context
	MonsterId string
} {
	var calls []struct {
		Ctx       context.Context
		MonsterId string
	}
	mock.lockRemove.RLock()
	calls = mock.calls.Remove
	mock.lockRemove.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *MonsterRepositoryMock) Save(ctx context.Context, m *Monster) error {
	if mock.SaveFunc == nil {
		panic("MonsterRepositoryMock.SaveFunc: method is nil but MonsterRepository.Save was just called")
	}
	callInfo := struct {
		Ctx context.Context
		M   *Monster
	}{
		Ctx: ctx,
		M:   m,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(ctx, m)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//
//	len(mockedMonsterRepository.SaveCalls())
func (mock *MonsterRepositoryMock) SaveCalls() []struct {
	Ctx context.Context
	M   *Monster
} {
	var calls []struct {
		Ctx context.Context
		M   *Monster
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}
