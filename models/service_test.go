package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/services-scrapper/common/db"
	"github.com/webmalc/services-scrapper/common/test"
)

// Should migrate the DB.
func TestMigrate(t *testing.T) {
	am := &AutoMigraterMock{}
	conn := db.NewConnection()
	defer conn.Close()
	args := []interface{}{&Service{}, &Phone{}, &Email{}, &Image{}, &Link{}}
	am.On("AutoMigrate", args...).Return(conn.DB).Once()
	Migrate(am)
	am.AssertExpectations(t)
}

// Should Validate the Service object
func TestService_Validate(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	repo := NewServiceRepository(conn)
	s := repo.NewService("invalid URL", "kijiji", "invalid")
	errors := conn.Create(s).GetErrors()
	assert.NotEmpty(t, errors)
	s = repo.NewService("http://example.com/12", "invalid", "invalid")
	errors = conn.Create(s).GetErrors()
	assert.NotEmpty(t, errors)
	s = repo.NewService("http://example.com/12", "kijiji", "valid")
	errors = conn.Create(s).GetErrors()
	assert.Empty(t, errors)
	var count int
	conn.Find(&Service{}).Count(&count)
	assert.Equal(t, 1, count)
}

// Should run the hook
func TestService_BeforeSave(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	repo := NewServiceRepository(conn)
	s := repo.NewService("http://example.com/12", "kijiji", "title")
	m := &ServiceProcessorMock{}
	s.Processor = m
	m.On("ProcessName", s).Return().Once()
	m.On("GatherEmails", s).Return().Once()
	m.On("GatherPhones", s).Return().Once()
	errors := conn.Create(s).GetErrors()
	assert.Empty(t, errors)
	m.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
