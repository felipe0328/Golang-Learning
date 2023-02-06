// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	models "github.com/golangLearning/ginSQL/controllers/employees/models"
	mock "github.com/stretchr/testify/mock"
)

// IEmployeeController is an autogenerated mock type for the IEmployeeController type
type IEmployeeController struct {
	mock.Mock
}

// CreateEmployee provides a mock function with given fields: employeeData
func (_m *IEmployeeController) CreateEmployee(employeeData models.Employee) (models.Employee, error) {
	ret := _m.Called(employeeData)

	var r0 models.Employee
	if rf, ok := ret.Get(0).(func(models.Employee) models.Employee); ok {
		r0 = rf(employeeData)
	} else {
		r0 = ret.Get(0).(models.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Employee) error); ok {
		r1 = rf(employeeData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEmployee provides a mock function with given fields: employeeId
func (_m *IEmployeeController) DeleteEmployee(employeeId int) error {
	ret := _m.Called(employeeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(employeeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEmployee provides a mock function with given fields: employeeId
func (_m *IEmployeeController) GetEmployee(employeeId int) (models.Employee, error) {
	ret := _m.Called(employeeId)

	var r0 models.Employee
	if rf, ok := ret.Get(0).(func(int) models.Employee); ok {
		r0 = rf(employeeId)
	} else {
		r0 = ret.Get(0).(models.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(employeeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmployeesList provides a mock function with given fields:
func (_m *IEmployeeController) GetEmployeesList() ([]models.Employee, error) {
	ret := _m.Called()

	var r0 []models.Employee
	if rf, ok := ret.Get(0).(func() []models.Employee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Employee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEmployee provides a mock function with given fields: employeeID, employeeData
func (_m *IEmployeeController) UpdateEmployee(employeeID int, employeeData models.Employee) (models.Employee, error) {
	ret := _m.Called(employeeID, employeeData)

	var r0 models.Employee
	if rf, ok := ret.Get(0).(func(int, models.Employee) models.Employee); ok {
		r0 = rf(employeeID, employeeData)
	} else {
		r0 = ret.Get(0).(models.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, models.Employee) error); ok {
		r1 = rf(employeeID, employeeData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIEmployeeController interface {
	mock.TestingT
	Cleanup(func())
}

// NewIEmployeeController creates a new instance of IEmployeeController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIEmployeeController(t mockConstructorTestingTNewIEmployeeController) *IEmployeeController {
	mock := &IEmployeeController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
