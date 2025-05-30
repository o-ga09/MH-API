// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package armors

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
//			GetAllFunc: func(ctx context.Context) (Armors, error) {
//				panic("mock out the GetAll method")
//			},
//			GetByIDFunc: func(ctx context.Context, armorId string) (*Armor, error) {
//				panic("mock out the GetByID method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context) (Armors, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, armorId string) (*Armor, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ArmorId is the armorId argument value.
			ArmorId string
		}
	}
	lockGetAll  sync.RWMutex
	lockGetByID sync.RWMutex
}

// GetAll calls GetAllFunc.
func (mock *RepositoryMock) GetAll(ctx context.Context) (Armors, error) {
	if mock.GetAllFunc == nil {
		panic("RepositoryMock.GetAllFunc: method is nil but Repository.GetAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(ctx)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//
//	len(mockedRepository.GetAllCalls())
func (mock *RepositoryMock) GetAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *RepositoryMock) GetByID(ctx context.Context, armorId string) (*Armor, error) {
	if mock.GetByIDFunc == nil {
		panic("RepositoryMock.GetByIDFunc: method is nil but Repository.GetByID was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		ArmorId string
	}{
		Ctx:     ctx,
		ArmorId: armorId,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, armorId)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//
//	len(mockedRepository.GetByIDCalls())
func (mock *RepositoryMock) GetByIDCalls() []struct {
	Ctx     context.Context
	ArmorId string
} {
	var calls []struct {
		Ctx     context.Context
		ArmorId string
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}
