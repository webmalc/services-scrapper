package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/services-scrapper/common/db"
	"github.com/webmalc/services-scrapper/common/test"
	"github.com/webmalc/services-scrapper/models/mocks"
)

// Should migrate the DB.
func TestMigrate(t *testing.T) {
	am := &mocks.AutoMigrater{}
	conn := db.NewConnection()
	am.On("AutoMigrate", mock.Anything).Return(conn.DB).Once()
	Migrate(am)
	am.AssertExpectations(t)
}

// Should create the Service object
func TestCreate(t *testing.T) {
	conn := db.NewConnection()
	Migrate(conn)
	errors := conn.Create(&Service{URL: "invalid URL"}).GetErrors()
	assert.NotEmpty(t, errors)
	errors = conn.Create(&Service{URL: "http://example.com/12"}).GetErrors()
	assert.Empty(t, errors)
	var count int
	conn.Find(&Service{}).Count(&count)
	assert.Equal(t, 1, count)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
