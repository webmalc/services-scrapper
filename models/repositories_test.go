package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/services-scrapper/common/db"
)

func TestServiceRepository_NewService(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	repo := NewServiceRepository(conn)
	s := repo.NewService("http://example.com", "kijiji", "title")
	assert.Equal(t, "title", s.Title)
	assert.Equal(t, "kijiji", s.Source)
	assert.Equal(t, "http://example.com", s.URL)
}

// Should add the first name and last name to the service
func TestServiceRepository_ProcessName(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	repo := NewServiceRepository(conn)
	s := repo.NewService("http://example.com", "kijiji", "title")
	s.FullName = "Anonymous"
	repo.ProcessName(s)
	assert.Equal(t, "", s.FirstName)
	assert.Equal(t, "", s.LastName)

	s.FullName = "Linus Torvalds"
	repo.ProcessName(s)
	assert.Equal(t, "Linus", s.FirstName)
	assert.Equal(t, "Torvalds", s.LastName)

	s.FullName = "Guido van Rossum"
	repo.ProcessName(s)
	assert.Equal(t, "Guido", s.FirstName)
	assert.Equal(t, "van", s.MiddleName)
	assert.Equal(t, "Rossum", s.LastName)
}

// Should add emails to the service
func TestServiceRepository_GatherEmails(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	repo := NewServiceRepository(conn)
	s := repo.NewService("http://example.com", "kijiji", "title")
	s.Emails = []Email{{Email: "zero@example.com"}}
	s.Description = "some text with emails@ one@example.com\ntwo@example.com"
	repo.GatherEmails(s)
	assert.Len(t, s.Emails, 3)
	assert.Equal(t, "zero@example.com", s.Emails[0].Email)
	assert.Equal(t, "one@example.com", s.Emails[1].Email)
	assert.Equal(t, "two@example.com", s.Emails[2].Email)
}

// Should add phones to the service
func TestServiceRepository_GatherPhones(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	repo := NewServiceRepository(conn)
	s := repo.NewService("http://example.com", "kijiji", "title")
	s.Phones = []Phone{{Phone: "89933443344"}}
	s.Description = "with 555-8909 phones123323 (333) 456 7899\ndd(111).456.7899"
	repo.GatherPhones(s)
	assert.Len(t, s.Phones, 4)
	assert.Equal(t, "89933443344", s.Phones[0].Phone)
	assert.Equal(t, "555-8909", s.Phones[1].Phone)
	assert.Equal(t, "123323 (333) 456 7899", s.Phones[2].Phone)
	assert.Equal(t, "(111).456.7899", s.Phones[3].Phone)
}

// Should return a new ServiceRepository
func TestNewServiceRepository(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	repo := NewServiceRepository(conn)
	assert.Equal(t, conn, repo.db)
}

// Should create a new service
func TestServiceRepository_CreateService(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	repo := NewServiceRepository(conn)
	s := repo.NewService("http://example.com", "kijiji", "title")
	s.FullName = "Tim Berners Lee"
	s.Description = "some text ne@example.com\n(333) 456 7899"
	errors := repo.CreateService(s)
	assert.Empty(t, errors)
	s = &Service{}
	conn.Set("gorm:auto_preload", true).First(s)
	assert.Equal(t, "Tim", s.FirstName)
	assert.Equal(t, "Berners", s.MiddleName)
	assert.Equal(t, "Lee", s.LastName)
	assert.Len(t, s.Phones, 1)
	assert.Len(t, s.Emails, 1)

	s = repo.NewService("http://example.com", "invalid", "title")
	errors = repo.CreateService(s)
	assert.NotEmpty(t, errors)
}

// Should check if a service with the provided URL exists
func TestServiceRepository_IsServiceWithURLExists(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	repo := NewServiceRepository(conn)
	s := repo.NewService("http://unique.com", "kijiji", "title")
	repo.CreateService(s)
	assert.True(t, repo.IsServiceWithURLExists("http://unique.com"))
	assert.False(t, repo.IsServiceWithURLExists("http://example.com"))
}

// Should append an image to a service
func TestServiceRepository_AppendImage(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	repo := NewServiceRepository(conn)
	s := repo.NewService("http://example.com", "kijiji", "title")
	repo.AppendImage("test image 1", s)
	repo.AppendImage("test image 2", s)
	assert.Equal(t, "test image 1", s.Images[0].Src)
	assert.Equal(t, "test image 2", s.Images[1].Src)
}
