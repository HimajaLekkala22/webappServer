// Code generated by mockery v2.42.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	request "github.com/mattermost/mattermost/server/public/shared/request"
	mock "github.com/stretchr/testify/mock"
)

// OutgoingOAuthConnectionInterface is an autogenerated mock type for the OutgoingOAuthConnectionInterface type
type OutgoingOAuthConnectionInterface struct {
	mock.Mock
}

// DeleteConnection provides a mock function with given fields: rctx, id
func (_m *OutgoingOAuthConnectionInterface) DeleteConnection(rctx request.CTX, id string) *model.AppError {
	ret := _m.Called(rctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteConnection")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.AppError); ok {
		r0 = rf(rctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// GetConnection provides a mock function with given fields: rctx, id
func (_m *OutgoingOAuthConnectionInterface) GetConnection(rctx request.CTX, id string) (*model.OutgoingOAuthConnection, *model.AppError) {
	ret := _m.Called(rctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetConnection")
	}

	var r0 *model.OutgoingOAuthConnection
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string) (*model.OutgoingOAuthConnection, *model.AppError)); ok {
		return rf(rctx, id)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.OutgoingOAuthConnection); ok {
		r0 = rf(rctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OutgoingOAuthConnection)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string) *model.AppError); ok {
		r1 = rf(rctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetConnectionForAudience provides a mock function with given fields: rctx, url
func (_m *OutgoingOAuthConnectionInterface) GetConnectionForAudience(rctx request.CTX, url string) (*model.OutgoingOAuthConnection, *model.AppError) {
	ret := _m.Called(rctx, url)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionForAudience")
	}

	var r0 *model.OutgoingOAuthConnection
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string) (*model.OutgoingOAuthConnection, *model.AppError)); ok {
		return rf(rctx, url)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.OutgoingOAuthConnection); ok {
		r0 = rf(rctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OutgoingOAuthConnection)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string) *model.AppError); ok {
		r1 = rf(rctx, url)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetConnections provides a mock function with given fields: rctx, filters
func (_m *OutgoingOAuthConnectionInterface) GetConnections(rctx request.CTX, filters model.OutgoingOAuthConnectionGetConnectionsFilter) ([]*model.OutgoingOAuthConnection, *model.AppError) {
	ret := _m.Called(rctx, filters)

	if len(ret) == 0 {
		panic("no return value specified for GetConnections")
	}

	var r0 []*model.OutgoingOAuthConnection
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, model.OutgoingOAuthConnectionGetConnectionsFilter) ([]*model.OutgoingOAuthConnection, *model.AppError)); ok {
		return rf(rctx, filters)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, model.OutgoingOAuthConnectionGetConnectionsFilter) []*model.OutgoingOAuthConnection); ok {
		r0 = rf(rctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OutgoingOAuthConnection)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, model.OutgoingOAuthConnectionGetConnectionsFilter) *model.AppError); ok {
		r1 = rf(rctx, filters)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// RetrieveTokenForConnection provides a mock function with given fields: rctx, conn
func (_m *OutgoingOAuthConnectionInterface) RetrieveTokenForConnection(rctx request.CTX, conn *model.OutgoingOAuthConnection) (*model.OutgoingOAuthConnectionToken, *model.AppError) {
	ret := _m.Called(rctx, conn)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveTokenForConnection")
	}

	var r0 *model.OutgoingOAuthConnectionToken
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.OutgoingOAuthConnection) (*model.OutgoingOAuthConnectionToken, *model.AppError)); ok {
		return rf(rctx, conn)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.OutgoingOAuthConnection) *model.OutgoingOAuthConnectionToken); ok {
		r0 = rf(rctx, conn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OutgoingOAuthConnectionToken)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.OutgoingOAuthConnection) *model.AppError); ok {
		r1 = rf(rctx, conn)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SanitizeConnection provides a mock function with given fields: conn
func (_m *OutgoingOAuthConnectionInterface) SanitizeConnection(conn *model.OutgoingOAuthConnection) {
	_m.Called(conn)
}

// SanitizeConnections provides a mock function with given fields: conns
func (_m *OutgoingOAuthConnectionInterface) SanitizeConnections(conns []*model.OutgoingOAuthConnection) {
	_m.Called(conns)
}

// SaveConnection provides a mock function with given fields: rctx, conn
func (_m *OutgoingOAuthConnectionInterface) SaveConnection(rctx request.CTX, conn *model.OutgoingOAuthConnection) (*model.OutgoingOAuthConnection, *model.AppError) {
	ret := _m.Called(rctx, conn)

	if len(ret) == 0 {
		panic("no return value specified for SaveConnection")
	}

	var r0 *model.OutgoingOAuthConnection
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.OutgoingOAuthConnection) (*model.OutgoingOAuthConnection, *model.AppError)); ok {
		return rf(rctx, conn)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.OutgoingOAuthConnection) *model.OutgoingOAuthConnection); ok {
		r0 = rf(rctx, conn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OutgoingOAuthConnection)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.OutgoingOAuthConnection) *model.AppError); ok {
		r1 = rf(rctx, conn)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpdateConnection provides a mock function with given fields: rctx, conn
func (_m *OutgoingOAuthConnectionInterface) UpdateConnection(rctx request.CTX, conn *model.OutgoingOAuthConnection) (*model.OutgoingOAuthConnection, *model.AppError) {
	ret := _m.Called(rctx, conn)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConnection")
	}

	var r0 *model.OutgoingOAuthConnection
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.OutgoingOAuthConnection) (*model.OutgoingOAuthConnection, *model.AppError)); ok {
		return rf(rctx, conn)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.OutgoingOAuthConnection) *model.OutgoingOAuthConnection); ok {
		r0 = rf(rctx, conn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OutgoingOAuthConnection)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.OutgoingOAuthConnection) *model.AppError); ok {
		r1 = rf(rctx, conn)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// NewOutgoingOAuthConnectionInterface creates a new instance of OutgoingOAuthConnectionInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOutgoingOAuthConnectionInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *OutgoingOAuthConnectionInterface {
	mock := &OutgoingOAuthConnectionInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}