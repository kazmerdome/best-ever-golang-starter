// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
	postQuerier "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post/post-querier"
)

// PostQuerier is an autogenerated mock type for the Querier type
type PostQuerier struct {
	mock.Mock
}

type PostQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *PostQuerier) EXPECT() *PostQuerier_Expecter {
	return &PostQuerier_Expecter{mock: &_m.Mock}
}

// CreateOne provides a mock function with given fields: ctx, arg
func (_m *PostQuerier) CreateOne(ctx context.Context, arg postQuerier.CreateOneParams) (postQuerier.Post, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateOne")
	}

	var r0 postQuerier.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, postQuerier.CreateOneParams) (postQuerier.Post, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, postQuerier.CreateOneParams) postQuerier.Post); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(postQuerier.Post)
	}

	if rf, ok := ret.Get(1).(func(context.Context, postQuerier.CreateOneParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostQuerier_CreateOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOne'
type PostQuerier_CreateOne_Call struct {
	*mock.Call
}

// CreateOne is a helper method to define mock.On call
//   - ctx context.Context
//   - arg postQuerier.CreateOneParams
func (_e *PostQuerier_Expecter) CreateOne(ctx interface{}, arg interface{}) *PostQuerier_CreateOne_Call {
	return &PostQuerier_CreateOne_Call{Call: _e.mock.On("CreateOne", ctx, arg)}
}

func (_c *PostQuerier_CreateOne_Call) Run(run func(ctx context.Context, arg postQuerier.CreateOneParams)) *PostQuerier_CreateOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(postQuerier.CreateOneParams))
	})
	return _c
}

func (_c *PostQuerier_CreateOne_Call) Return(_a0 postQuerier.Post, _a1 error) *PostQuerier_CreateOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostQuerier_CreateOne_Call) RunAndReturn(run func(context.Context, postQuerier.CreateOneParams) (postQuerier.Post, error)) *PostQuerier_CreateOne_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteOne provides a mock function with given fields: ctx, id
func (_m *PostQuerier) DeleteOne(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostQuerier_DeleteOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteOne'
type PostQuerier_DeleteOne_Call struct {
	*mock.Call
}

// DeleteOne is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *PostQuerier_Expecter) DeleteOne(ctx interface{}, id interface{}) *PostQuerier_DeleteOne_Call {
	return &PostQuerier_DeleteOne_Call{Call: _e.mock.On("DeleteOne", ctx, id)}
}

func (_c *PostQuerier_DeleteOne_Call) Run(run func(ctx context.Context, id uuid.UUID)) *PostQuerier_DeleteOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *PostQuerier_DeleteOne_Call) Return(_a0 error) *PostQuerier_DeleteOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostQuerier_DeleteOne_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *PostQuerier_DeleteOne_Call {
	_c.Call.Return(run)
	return _c
}

// GetMany provides a mock function with given fields: ctx, arg
func (_m *PostQuerier) GetMany(ctx context.Context, arg postQuerier.GetManyParams) ([]postQuerier.Post, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetMany")
	}

	var r0 []postQuerier.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, postQuerier.GetManyParams) ([]postQuerier.Post, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, postQuerier.GetManyParams) []postQuerier.Post); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]postQuerier.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, postQuerier.GetManyParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostQuerier_GetMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMany'
type PostQuerier_GetMany_Call struct {
	*mock.Call
}

// GetMany is a helper method to define mock.On call
//   - ctx context.Context
//   - arg postQuerier.GetManyParams
func (_e *PostQuerier_Expecter) GetMany(ctx interface{}, arg interface{}) *PostQuerier_GetMany_Call {
	return &PostQuerier_GetMany_Call{Call: _e.mock.On("GetMany", ctx, arg)}
}

func (_c *PostQuerier_GetMany_Call) Run(run func(ctx context.Context, arg postQuerier.GetManyParams)) *PostQuerier_GetMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(postQuerier.GetManyParams))
	})
	return _c
}

func (_c *PostQuerier_GetMany_Call) Return(_a0 []postQuerier.Post, _a1 error) *PostQuerier_GetMany_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostQuerier_GetMany_Call) RunAndReturn(run func(context.Context, postQuerier.GetManyParams) ([]postQuerier.Post, error)) *PostQuerier_GetMany_Call {
	_c.Call.Return(run)
	return _c
}

// GetManyByIds provides a mock function with given fields: ctx, dollar_1
func (_m *PostQuerier) GetManyByIds(ctx context.Context, dollar_1 []uuid.UUID) ([]postQuerier.Post, error) {
	ret := _m.Called(ctx, dollar_1)

	if len(ret) == 0 {
		panic("no return value specified for GetManyByIds")
	}

	var r0 []postQuerier.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) ([]postQuerier.Post, error)); ok {
		return rf(ctx, dollar_1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) []postQuerier.Post); ok {
		r0 = rf(ctx, dollar_1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]postQuerier.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uuid.UUID) error); ok {
		r1 = rf(ctx, dollar_1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostQuerier_GetManyByIds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetManyByIds'
type PostQuerier_GetManyByIds_Call struct {
	*mock.Call
}

// GetManyByIds is a helper method to define mock.On call
//   - ctx context.Context
//   - dollar_1 []uuid.UUID
func (_e *PostQuerier_Expecter) GetManyByIds(ctx interface{}, dollar_1 interface{}) *PostQuerier_GetManyByIds_Call {
	return &PostQuerier_GetManyByIds_Call{Call: _e.mock.On("GetManyByIds", ctx, dollar_1)}
}

func (_c *PostQuerier_GetManyByIds_Call) Run(run func(ctx context.Context, dollar_1 []uuid.UUID)) *PostQuerier_GetManyByIds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uuid.UUID))
	})
	return _c
}

func (_c *PostQuerier_GetManyByIds_Call) Return(_a0 []postQuerier.Post, _a1 error) *PostQuerier_GetManyByIds_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostQuerier_GetManyByIds_Call) RunAndReturn(run func(context.Context, []uuid.UUID) ([]postQuerier.Post, error)) *PostQuerier_GetManyByIds_Call {
	_c.Call.Return(run)
	return _c
}

// GetOneById provides a mock function with given fields: ctx, id
func (_m *PostQuerier) GetOneById(ctx context.Context, id uuid.UUID) (postQuerier.Post, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetOneById")
	}

	var r0 postQuerier.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (postQuerier.Post, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) postQuerier.Post); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(postQuerier.Post)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostQuerier_GetOneById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOneById'
type PostQuerier_GetOneById_Call struct {
	*mock.Call
}

// GetOneById is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *PostQuerier_Expecter) GetOneById(ctx interface{}, id interface{}) *PostQuerier_GetOneById_Call {
	return &PostQuerier_GetOneById_Call{Call: _e.mock.On("GetOneById", ctx, id)}
}

func (_c *PostQuerier_GetOneById_Call) Run(run func(ctx context.Context, id uuid.UUID)) *PostQuerier_GetOneById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *PostQuerier_GetOneById_Call) Return(_a0 postQuerier.Post, _a1 error) *PostQuerier_GetOneById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostQuerier_GetOneById_Call) RunAndReturn(run func(context.Context, uuid.UUID) (postQuerier.Post, error)) *PostQuerier_GetOneById_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOneById provides a mock function with given fields: ctx, arg
func (_m *PostQuerier) UpdateOneById(ctx context.Context, arg postQuerier.UpdateOneByIdParams) (postQuerier.Post, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOneById")
	}

	var r0 postQuerier.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, postQuerier.UpdateOneByIdParams) (postQuerier.Post, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, postQuerier.UpdateOneByIdParams) postQuerier.Post); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(postQuerier.Post)
	}

	if rf, ok := ret.Get(1).(func(context.Context, postQuerier.UpdateOneByIdParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostQuerier_UpdateOneById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOneById'
type PostQuerier_UpdateOneById_Call struct {
	*mock.Call
}

// UpdateOneById is a helper method to define mock.On call
//   - ctx context.Context
//   - arg postQuerier.UpdateOneByIdParams
func (_e *PostQuerier_Expecter) UpdateOneById(ctx interface{}, arg interface{}) *PostQuerier_UpdateOneById_Call {
	return &PostQuerier_UpdateOneById_Call{Call: _e.mock.On("UpdateOneById", ctx, arg)}
}

func (_c *PostQuerier_UpdateOneById_Call) Run(run func(ctx context.Context, arg postQuerier.UpdateOneByIdParams)) *PostQuerier_UpdateOneById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(postQuerier.UpdateOneByIdParams))
	})
	return _c
}

func (_c *PostQuerier_UpdateOneById_Call) Return(_a0 postQuerier.Post, _a1 error) *PostQuerier_UpdateOneById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostQuerier_UpdateOneById_Call) RunAndReturn(run func(context.Context, postQuerier.UpdateOneByIdParams) (postQuerier.Post, error)) *PostQuerier_UpdateOneById_Call {
	_c.Call.Return(run)
	return _c
}

// NewPostQuerier creates a new instance of PostQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostQuerier {
	mock := &PostQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}