// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/api/config/models"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetConfig provides a mock function with given fields: params
func (_m *Usecase) GetConfig(params *models.ConfigParams) *models.ConfigBag {
	ret := _m.Called(params)

	var r0 *models.ConfigBag
	if rf, ok := ret.Get(0).(func(*models.ConfigParams) *models.ConfigBag); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ConfigBag)
		}
	}

	return r0
}

// Hydrate provides a mock function with given fields: _a0
func (_m *Usecase) Hydrate(_a0 *models.ConfigBag) {
	_m.Called(_a0)
}

// Verify provides a mock function with given fields: _a0
func (_m *Usecase) Verify(_a0 *models.ConfigBag) {
	_m.Called(_a0)
}
