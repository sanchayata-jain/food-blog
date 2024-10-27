// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/sanchayata-jain/food-blog/internal/recipes/models"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateRecipe provides a mock function with given fields: ctx, recipe
func (_m *Service) CreateRecipe(ctx context.Context, recipe *models.Recipe) error {
	ret := _m.Called(ctx, recipe)

	if len(ret) == 0 {
		panic("no return value specified for CreateRecipe")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Recipe) error); ok {
		r0 = rf(ctx, recipe)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
