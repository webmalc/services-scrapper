package models

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

// ServiceProcessorMock is a mock struct
type ServiceProcessorMock struct {
	mock.Mock
}

// GatherEmails is a method mock
func (s *ServiceProcessorMock) GatherEmails(service *Service) {
	s.Called(service)
}

// GatherPhones is a method mock
func (s *ServiceProcessorMock) GatherPhones(service *Service) {
	s.Called(service)
}

// ProcessName is a method mock
func (s *ServiceProcessorMock) ProcessName(service *Service) {
	s.Called(service)
}

// AutoMigraterMock logs errors.
type AutoMigraterMock struct {
	mock.Mock
}

// AutoMigrate is a method mock
func (m *AutoMigraterMock) AutoMigrate(values ...interface{}) *gorm.DB {
	arg := m.Called(values...)
	return arg.Get(0).(*gorm.DB)
}
