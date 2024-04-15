// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	category "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CategoryService is an autogenerated mock type for the CategoryService type
type CategoryService struct {
	mock.Mock
}

type CategoryService_Expecter struct {
	mock *mock.Mock
}

func (_m *CategoryService) EXPECT() *CategoryService_Expecter {
	return &CategoryService_Expecter{mock: &_m.Mock}
}

// CreateCategory provides a mock function with given fields: ctx, data
func (_m *CategoryService) CreateCategory(ctx context.Context, data category.CreateDto) (*category.Category, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for CreateCategory")
	}

	var r0 *category.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, category.CreateDto) (*category.Category, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, category.CreateDto) *category.Category); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*category.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, category.CreateDto) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryService_CreateCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCategory'
type CategoryService_CreateCategory_Call struct {
	*mock.Call
}

// CreateCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - data category.CreateDto
func (_e *CategoryService_Expecter) CreateCategory(ctx interface{}, data interface{}) *CategoryService_CreateCategory_Call {
	return &CategoryService_CreateCategory_Call{Call: _e.mock.On("CreateCategory", ctx, data)}
}

func (_c *CategoryService_CreateCategory_Call) Run(run func(ctx context.Context, data category.CreateDto)) *CategoryService_CreateCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(category.CreateDto))
	})
	return _c
}

func (_c *CategoryService_CreateCategory_Call) Return(_a0 *category.Category, _a1 error) *CategoryService_CreateCategory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryService_CreateCategory_Call) RunAndReturn(run func(context.Context, category.CreateDto) (*category.Category, error)) *CategoryService_CreateCategory_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCategory provides a mock function with given fields: ctx, id
func (_m *CategoryService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCategory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CategoryService_DeleteCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCategory'
type CategoryService_DeleteCategory_Call struct {
	*mock.Call
}

// DeleteCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *CategoryService_Expecter) DeleteCategory(ctx interface{}, id interface{}) *CategoryService_DeleteCategory_Call {
	return &CategoryService_DeleteCategory_Call{Call: _e.mock.On("DeleteCategory", ctx, id)}
}

func (_c *CategoryService_DeleteCategory_Call) Run(run func(ctx context.Context, id uuid.UUID)) *CategoryService_DeleteCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *CategoryService_DeleteCategory_Call) Return(_a0 error) *CategoryService_DeleteCategory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CategoryService_DeleteCategory_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *CategoryService_DeleteCategory_Call {
	_c.Call.Return(run)
	return _c
}

// GetCategory provides a mock function with given fields: ctx, id
func (_m *CategoryService) GetCategory(ctx context.Context, id uuid.UUID) (*category.Category, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetCategory")
	}

	var r0 *category.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*category.Category, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *category.Category); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*category.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryService_GetCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCategory'
type CategoryService_GetCategory_Call struct {
	*mock.Call
}

// GetCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *CategoryService_Expecter) GetCategory(ctx interface{}, id interface{}) *CategoryService_GetCategory_Call {
	return &CategoryService_GetCategory_Call{Call: _e.mock.On("GetCategory", ctx, id)}
}

func (_c *CategoryService_GetCategory_Call) Run(run func(ctx context.Context, id uuid.UUID)) *CategoryService_GetCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *CategoryService_GetCategory_Call) Return(_a0 *category.Category, _a1 error) *CategoryService_GetCategory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryService_GetCategory_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*category.Category, error)) *CategoryService_GetCategory_Call {
	_c.Call.Return(run)
	return _c
}

// ListCategories provides a mock function with given fields: ctx, where
func (_m *CategoryService) ListCategories(ctx context.Context, where *category.WhereDto) ([]category.Category, error) {
	ret := _m.Called(ctx, where)

	if len(ret) == 0 {
		panic("no return value specified for ListCategories")
	}

	var r0 []category.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *category.WhereDto) ([]category.Category, error)); ok {
		return rf(ctx, where)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *category.WhereDto) []category.Category); ok {
		r0 = rf(ctx, where)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]category.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *category.WhereDto) error); ok {
		r1 = rf(ctx, where)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryService_ListCategories_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCategories'
type CategoryService_ListCategories_Call struct {
	*mock.Call
}

// ListCategories is a helper method to define mock.On call
//   - ctx context.Context
//   - where *category.WhereDto
func (_e *CategoryService_Expecter) ListCategories(ctx interface{}, where interface{}) *CategoryService_ListCategories_Call {
	return &CategoryService_ListCategories_Call{Call: _e.mock.On("ListCategories", ctx, where)}
}

func (_c *CategoryService_ListCategories_Call) Run(run func(ctx context.Context, where *category.WhereDto)) *CategoryService_ListCategories_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*category.WhereDto))
	})
	return _c
}

func (_c *CategoryService_ListCategories_Call) Return(_a0 []category.Category, _a1 error) *CategoryService_ListCategories_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryService_ListCategories_Call) RunAndReturn(run func(context.Context, *category.WhereDto) ([]category.Category, error)) *CategoryService_ListCategories_Call {
	_c.Call.Return(run)
	return _c
}

// LoadCategory provides a mock function with given fields: ctx, id
func (_m *CategoryService) LoadCategory(ctx context.Context, id uuid.UUID) (*category.Category, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for LoadCategory")
	}

	var r0 *category.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*category.Category, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *category.Category); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*category.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryService_LoadCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadCategory'
type CategoryService_LoadCategory_Call struct {
	*mock.Call
}

// LoadCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *CategoryService_Expecter) LoadCategory(ctx interface{}, id interface{}) *CategoryService_LoadCategory_Call {
	return &CategoryService_LoadCategory_Call{Call: _e.mock.On("LoadCategory", ctx, id)}
}

func (_c *CategoryService_LoadCategory_Call) Run(run func(ctx context.Context, id uuid.UUID)) *CategoryService_LoadCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *CategoryService_LoadCategory_Call) Return(_a0 *category.Category, _a1 error) *CategoryService_LoadCategory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryService_LoadCategory_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*category.Category, error)) *CategoryService_LoadCategory_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCategory provides a mock function with given fields: ctx, id, data
func (_m *CategoryService) UpdateCategory(ctx context.Context, id uuid.UUID, data category.UpdateDto) (*category.Category, error) {
	ret := _m.Called(ctx, id, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCategory")
	}

	var r0 *category.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, category.UpdateDto) (*category.Category, error)); ok {
		return rf(ctx, id, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, category.UpdateDto) *category.Category); ok {
		r0 = rf(ctx, id, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*category.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, category.UpdateDto) error); ok {
		r1 = rf(ctx, id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryService_UpdateCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCategory'
type CategoryService_UpdateCategory_Call struct {
	*mock.Call
}

// UpdateCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
//   - data category.UpdateDto
func (_e *CategoryService_Expecter) UpdateCategory(ctx interface{}, id interface{}, data interface{}) *CategoryService_UpdateCategory_Call {
	return &CategoryService_UpdateCategory_Call{Call: _e.mock.On("UpdateCategory", ctx, id, data)}
}

func (_c *CategoryService_UpdateCategory_Call) Run(run func(ctx context.Context, id uuid.UUID, data category.UpdateDto)) *CategoryService_UpdateCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(category.UpdateDto))
	})
	return _c
}

func (_c *CategoryService_UpdateCategory_Call) Return(_a0 *category.Category, _a1 error) *CategoryService_UpdateCategory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryService_UpdateCategory_Call) RunAndReturn(run func(context.Context, uuid.UUID, category.UpdateDto) (*category.Category, error)) *CategoryService_UpdateCategory_Call {
	_c.Call.Return(run)
	return _c
}

// NewCategoryService creates a new instance of CategoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCategoryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CategoryService {
	mock := &CategoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}