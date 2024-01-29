// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	models "projects/LDmitryLD/parser/app/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// SQLAdapterer is an autogenerated mock type for the SQLAdapterer type
type SQLAdapterer struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *SQLAdapterer) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: vac
func (_m *SQLAdapterer) Insert(vac models.Vacancy) (int, error) {
	ret := _m.Called(vac)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Vacancy) (int, error)); ok {
		return rf(vac)
	}
	if rf, ok := ret.Get(0).(func(models.Vacancy) int); ok {
		r0 = rf(vac)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(models.Vacancy) error); ok {
		r1 = rf(vac)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectVacancies provides a mock function with given fields: query, list
func (_m *SQLAdapterer) SelectVacancies(query string, list bool) ([]models.Vacancy, error) {
	ret := _m.Called(query, list)

	var r0 []models.Vacancy
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) ([]models.Vacancy, error)); ok {
		return rf(query, list)
	}
	if rf, ok := ret.Get(0).(func(string, bool) []models.Vacancy); ok {
		r0 = rf(query, list)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Vacancy)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(query, list)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectVacancy provides a mock function with given fields: id
func (_m *SQLAdapterer) SelectVacancy(id int) (models.Vacancy, error) {
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

// NewSQLAdapterer creates a new instance of SQLAdapterer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSQLAdapterer(t interface {
	mock.TestingT
	Cleanup(func())
}) *SQLAdapterer {
	mock := &SQLAdapterer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
