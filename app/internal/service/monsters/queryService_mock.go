// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package monsters

import (
	"context"
	"sync"
)

// Ensure, that MonsterQueryServiceMock does implement MonsterQueryService.
// If this is not the case, regenerate this file with moq.
var _ MonsterQueryService = &MonsterQueryServiceMock{}

// MonsterQueryServiceMock is a mock implementation of MonsterQueryService.
//
//	func TestSomethingThatUsesMonsterQueryService(t *testing.T) {
//
//		// make and configure a mocked MonsterQueryService
//		mockedMonsterQueryService := &MonsterQueryServiceMock{
//			FetchListFunc: func(ctx context.Context, id string) ([]*FetchMonsterListDto, error) {
//				panic("mock out the FetchList method")
//			},
//			FetchRankFunc: func(ctx context.Context) ([]*FetchMonsterRankingDto, error) {
//				panic("mock out the FetchRank method")
//			},
//		}
//
//		// use mockedMonsterQueryService in code that requires MonsterQueryService
//		// and then make assertions.
//
//	}
type MonsterQueryServiceMock struct {
	// FetchListFunc mocks the FetchList method.
	FetchListFunc func(ctx context.Context, id string) ([]*FetchMonsterListDto, error)

	// FetchRankFunc mocks the FetchRank method.
	FetchRankFunc func(ctx context.Context) ([]*FetchMonsterRankingDto, error)

	// calls tracks calls to the methods.
	calls struct {
		// FetchList holds details about calls to the FetchList method.
		FetchList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// FetchRank holds details about calls to the FetchRank method.
		FetchRank []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockFetchList sync.RWMutex
	lockFetchRank sync.RWMutex
}

// FetchList calls FetchListFunc.
func (mock *MonsterQueryServiceMock) FetchList(ctx context.Context, id string) ([]*FetchMonsterListDto, error) {
	if mock.FetchListFunc == nil {
		panic("MonsterQueryServiceMock.FetchListFunc: method is nil but MonsterQueryService.FetchList was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFetchList.Lock()
	mock.calls.FetchList = append(mock.calls.FetchList, callInfo)
	mock.lockFetchList.Unlock()
	return mock.FetchListFunc(ctx, id)
}

// FetchListCalls gets all the calls that were made to FetchList.
// Check the length with:
//
//	len(mockedMonsterQueryService.FetchListCalls())
func (mock *MonsterQueryServiceMock) FetchListCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockFetchList.RLock()
	calls = mock.calls.FetchList
	mock.lockFetchList.RUnlock()
	return calls
}

// FetchRank calls FetchRankFunc.
func (mock *MonsterQueryServiceMock) FetchRank(ctx context.Context) ([]*FetchMonsterRankingDto, error) {
	if mock.FetchRankFunc == nil {
		panic("MonsterQueryServiceMock.FetchRankFunc: method is nil but MonsterQueryService.FetchRank was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockFetchRank.Lock()
	mock.calls.FetchRank = append(mock.calls.FetchRank, callInfo)
	mock.lockFetchRank.Unlock()
	return mock.FetchRankFunc(ctx)
}

// FetchRankCalls gets all the calls that were made to FetchRank.
// Check the length with:
//
//	len(mockedMonsterQueryService.FetchRankCalls())
func (mock *MonsterQueryServiceMock) FetchRankCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockFetchRank.RLock()
	calls = mock.calls.FetchRank
	mock.lockFetchRank.RUnlock()
	return calls
}