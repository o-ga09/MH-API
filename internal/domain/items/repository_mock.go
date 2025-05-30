// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package items

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
//			FindAllFunc: func(ctx context.Context) (Items, error) {
//				panic("mock out the FindAll method")
//			},
//			FindByIDFunc: func(ctx context.Context, itemID string) (*Item, error) {
//				panic("mock out the FindByID method")
//			},
//			FindByMonsterIDFunc: func(ctx context.Context, monsterID string) (Items, error) {
//				panic("mock out the FindByMonsterID method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// FindAllFunc mocks the FindAll method.
	FindAllFunc func(ctx context.Context) (Items, error)

	// FindByIDFunc mocks the FindByID method.
	FindByIDFunc func(ctx context.Context, itemID string) (*Item, error)

	// FindByMonsterIDFunc mocks the FindByMonsterID method.
	FindByMonsterIDFunc func(ctx context.Context, monsterID string) (Items, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindAll holds details about calls to the FindAll method.
		FindAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// FindByID holds details about calls to the FindByID method.
		FindByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ItemID is the itemID argument value.
			ItemID string
		}
		// FindByMonsterID holds details about calls to the FindByMonsterID method.
		FindByMonsterID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// MonsterID is the monsterID argument value.
			MonsterID string
		}
	}
	lockFindAll         sync.RWMutex
	lockFindByID        sync.RWMutex
	lockFindByMonsterID sync.RWMutex
}

// FindAll calls FindAllFunc.
func (mock *RepositoryMock) FindAll(ctx context.Context) (Items, error) {
	if mock.FindAllFunc == nil {
		panic("RepositoryMock.FindAllFunc: method is nil but Repository.FindAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockFindAll.Lock()
	mock.calls.FindAll = append(mock.calls.FindAll, callInfo)
	mock.lockFindAll.Unlock()
	return mock.FindAllFunc(ctx)
}

// FindAllCalls gets all the calls that were made to FindAll.
// Check the length with:
//
//	len(mockedRepository.FindAllCalls())
func (mock *RepositoryMock) FindAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockFindAll.RLock()
	calls = mock.calls.FindAll
	mock.lockFindAll.RUnlock()
	return calls
}

// FindByID calls FindByIDFunc.
func (mock *RepositoryMock) FindByID(ctx context.Context, itemID string) (*Item, error) {
	if mock.FindByIDFunc == nil {
		panic("RepositoryMock.FindByIDFunc: method is nil but Repository.FindByID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		ItemID string
	}{
		Ctx:    ctx,
		ItemID: itemID,
	}
	mock.lockFindByID.Lock()
	mock.calls.FindByID = append(mock.calls.FindByID, callInfo)
	mock.lockFindByID.Unlock()
	return mock.FindByIDFunc(ctx, itemID)
}

// FindByIDCalls gets all the calls that were made to FindByID.
// Check the length with:
//
//	len(mockedRepository.FindByIDCalls())
func (mock *RepositoryMock) FindByIDCalls() []struct {
	Ctx    context.Context
	ItemID string
} {
	var calls []struct {
		Ctx    context.Context
		ItemID string
	}
	mock.lockFindByID.RLock()
	calls = mock.calls.FindByID
	mock.lockFindByID.RUnlock()
	return calls
}

// FindByMonsterID calls FindByMonsterIDFunc.
func (mock *RepositoryMock) FindByMonsterID(ctx context.Context, monsterID string) (Items, error) {
	if mock.FindByMonsterIDFunc == nil {
		panic("RepositoryMock.FindByMonsterIDFunc: method is nil but Repository.FindByMonsterID was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		MonsterID string
	}{
		Ctx:       ctx,
		MonsterID: monsterID,
	}
	mock.lockFindByMonsterID.Lock()
	mock.calls.FindByMonsterID = append(mock.calls.FindByMonsterID, callInfo)
	mock.lockFindByMonsterID.Unlock()
	return mock.FindByMonsterIDFunc(ctx, monsterID)
}

// FindByMonsterIDCalls gets all the calls that were made to FindByMonsterID.
// Check the length with:
//
//	len(mockedRepository.FindByMonsterIDCalls())
func (mock *RepositoryMock) FindByMonsterIDCalls() []struct {
	Ctx       context.Context
	MonsterID string
} {
	var calls []struct {
		Ctx       context.Context
		MonsterID string
	}
	mock.lockFindByMonsterID.RLock()
	calls = mock.calls.FindByMonsterID
	mock.lockFindByMonsterID.RUnlock()
	return calls
}
