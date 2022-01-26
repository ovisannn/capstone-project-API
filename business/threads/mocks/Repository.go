// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	threads "disspace/business/threads"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *Repository) Create(ctx context.Context, data *threads.Domain) (threads.Domain, error) {
	ret := _m.Called(ctx, data)

	var r0 threads.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *threads.Domain) threads.Domain); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(threads.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *threads.Domain) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, sort
func (_m *Repository) GetAll(ctx context.Context, sort string) ([]threads.Domain, error) {
	ret := _m.Called(ctx, sort)

	var r0 []threads.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) []threads.Domain); ok {
		r0 = rf(ctx, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]threads.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id string) (threads.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 threads.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) threads.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(threads.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, q, sort
func (_m *Repository) Search(ctx context.Context, q string, sort string) ([]threads.Domain, error) {
	ret := _m.Called(ctx, q, sort)

	var r0 []threads.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []threads.Domain); ok {
		r0 = rf(ctx, q, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]threads.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, q, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, data, id
func (_m *Repository) Update(ctx context.Context, data *threads.Domain, id string) error {
	ret := _m.Called(ctx, data, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *threads.Domain, string) error); ok {
		r0 = rf(ctx, data, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}