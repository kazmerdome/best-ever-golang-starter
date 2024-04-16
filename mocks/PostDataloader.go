// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	post "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post"

	uuid "github.com/google/uuid"
)

// PostDataloader is an autogenerated mock type for the PostDataloader type
type PostDataloader struct {
	mock.Mock
}

type PostDataloader_Expecter struct {
	mock *mock.Mock
}

func (_m *PostDataloader) EXPECT() *PostDataloader_Expecter {
	return &PostDataloader_Expecter{mock: &_m.Mock}
}

// ItemLoader provides a mock function with given fields: ctx, id
func (_m *PostDataloader) ItemLoader(ctx context.Context, id uuid.UUID) (*post.Post, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for ItemLoader")
	}

	var r0 *post.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*post.Post, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *post.Post); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*post.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostDataloader_ItemLoader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ItemLoader'
type PostDataloader_ItemLoader_Call struct {
	*mock.Call
}

// ItemLoader is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *PostDataloader_Expecter) ItemLoader(ctx interface{}, id interface{}) *PostDataloader_ItemLoader_Call {
	return &PostDataloader_ItemLoader_Call{Call: _e.mock.On("ItemLoader", ctx, id)}
}

func (_c *PostDataloader_ItemLoader_Call) Run(run func(ctx context.Context, id uuid.UUID)) *PostDataloader_ItemLoader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *PostDataloader_ItemLoader_Call) Return(_a0 *post.Post, _a1 error) *PostDataloader_ItemLoader_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PostDataloader_ItemLoader_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*post.Post, error)) *PostDataloader_ItemLoader_Call {
	_c.Call.Return(run)
	return _c
}

// NewPostDataloader creates a new instance of PostDataloader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostDataloader(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostDataloader {
	mock := &PostDataloader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}