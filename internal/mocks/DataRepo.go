// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	repository "github.com/aswinudhayakumar/account-transactions/pkg/repository"
	mock "github.com/stretchr/testify/mock"
)

// DataRepo is an autogenerated mock type for the DataRepo type
type DataRepo struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, req
func (_m *DataRepo) CreateAccount(ctx context.Context, req repository.CreateAccountReqParams) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.CreateAccountReqParams) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccountByAccountID provides a mock function with given fields: ctx, accountID
func (_m *DataRepo) GetAccountByAccountID(ctx context.Context, accountID int) (*repository.AccountResponse, error) {
	ret := _m.Called(ctx, accountID)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountByAccountID")
	}

	var r0 *repository.AccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*repository.AccountResponse, error)); ok {
		return rf(ctx, accountID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *repository.AccountResponse); ok {
		r0 = rf(ctx, accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.AccountResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDataRepo creates a new instance of DataRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataRepo {
	mock := &DataRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
