// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package Products

import (
	"context"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked Repository
//		mockedRepository := &RepositoryMock{
//			GetFunc: func(ctx context.Context, monsterId string) (*Products, error) {
//				panic("mock out the Get method")
//			},
//			RemoveFunc: func(ctx context.Context, monsterId string) error {
//				panic("mock out the Remove method")
//			},
//			SaveFunc: func(ctx context.Context, m Product) error {
//				panic("mock out the Save method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, monsterId string) (*Products, error)

	// RemoveFunc mocks the Remove method.
	RemoveFunc func(ctx context.Context, monsterId string) error

	// SaveFunc mocks the Save method.
	SaveFunc func(ctx context.Context, m Product) error

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// MonsterId is the monsterId argument value.
			MonsterId string
		}
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
			M Product
		}
	}
	lockGet    sync.RWMutex
	lockRemove sync.RWMutex
	lockSave   sync.RWMutex
}

// Get calls GetFunc.
func (mock *RepositoryMock) Get(ctx context.Context, monsterId string) (*Products, error) {
	if mock.GetFunc == nil {
		panic("RepositoryMock.GetFunc: method is nil but Repository.Get was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		MonsterId string
	}{
		Ctx:       ctx,
		MonsterId: monsterId,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, monsterId)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedRepository.GetCalls())
func (mock *RepositoryMock) GetCalls() []struct {
	Ctx       context.Context
	MonsterId string
} {
	var calls []struct {
		Ctx       context.Context
		MonsterId string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// Remove calls RemoveFunc.
func (mock *RepositoryMock) Remove(ctx context.Context, monsterId string) error {
	if mock.RemoveFunc == nil {
		panic("RepositoryMock.RemoveFunc: method is nil but Repository.Remove was just called")
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
//	len(mockedRepository.RemoveCalls())
func (mock *RepositoryMock) RemoveCalls() []struct {
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
func (mock *RepositoryMock) Save(ctx context.Context, m Product) error {
	if mock.SaveFunc == nil {
		panic("RepositoryMock.SaveFunc: method is nil but Repository.Save was just called")
	}
	callInfo := struct {
		Ctx context.Context
		M   Product
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
//	len(mockedRepository.SaveCalls())
func (mock *RepositoryMock) SaveCalls() []struct {
	Ctx context.Context
	M   Product
} {
	var calls []struct {
		Ctx context.Context
		M   Product
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}
