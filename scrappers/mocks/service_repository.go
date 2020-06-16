package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/services-scrapper/models"
)

// ServiceRepository is a mock struct
type ServiceRepository struct {
	mock.Mock
}

// NewService is a method mock
func (s *ServiceRepository) NewService(
	url, source, title string,
) *models.Service {
	arg := s.Called(url, source, title)
	return arg.Get(0).(*models.Service)
}

// CreateService is a method mock
func (s *ServiceRepository) CreateService(
	service *models.Service,
) []error {
	arg := s.Called(service)
	return arg.Get(0).([]error)
}

// IsServiceWithURLExists is a method mock
func (s *ServiceRepository) IsServiceWithURLExists(url string) bool {
	arg := s.Called(url)
	return arg.Get(0).(bool)
}

// AppendImage is a method mock
func (s *ServiceRepository) AppendImage(src string, service *models.Service) {
	s.Called(src, service)
}
