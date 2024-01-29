// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	models "projects/LDmitryLD/parser/app/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Vacancyer is an autogenerated mock type for the Vacancyer type
type Vacancyer struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Vacancyer) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *Vacancyer) Get(id int) (models.Vacancy, error) {
	ret := _m.Called(id)

	var r0 models.Vacancy
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (models.Vacancy, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) models.Vacancy); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Vacancy)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *Vacancyer) List() ([]models.Vacancy, error) {
	ret := _m.Called()

	var r0 []models.Vacancy
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Vacancy, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Vacancy); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Vacancy)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: query
func (_m *Vacancyer) Search(query string) ([]models.Vacancy, error) {
	ret := _m.Called(query)

	var r0 []models.Vacancy
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.Vacancy, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(string) []models.Vacancy); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Vacancy)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewVacancyer creates a new instance of Vacancyer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVacancyer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Vacancyer {
	mock := &Vacancyer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}