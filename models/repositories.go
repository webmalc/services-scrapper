package models

import (
	"regexp"
	"strings"
)

// ServiceRepository is the service repository struct
type ServiceRepository struct {
	db Database
}

// NewService return a service object
func (r *ServiceRepository) NewService(url, source, title string) *Service {
	return &Service{
		URL:       url,
		Source:    source,
		Title:     title,
		Processor: r,
	}
}

// CreateService creates a new service
func (r *ServiceRepository) CreateService(service *Service) []error {
	errors := r.db.Create(service).GetErrors()
	return errors
}

// IsServiceWithURLExists checks if a service with the provided URL exists
func (r *ServiceRepository) IsServiceWithURLExists(url string) bool {
	var count int
	r.db.Model(&Service{}).Where("url = ?", url).Count(&count)
	return count > 0
}

// ProcessName tries to get the first name and last name from the full name
func (r *ServiceRepository) ProcessName(service *Service) {
	const withMiddlename = 3
	const noMiddleName = withMiddlename - 1
	parts := strings.Fields(service.FullName)
	switch len(parts) {
	case noMiddleName:
		service.FirstName = parts[0]
		service.LastName = parts[1]
	case withMiddlename:
		service.FirstName = parts[0]
		service.MiddleName = parts[1]
		service.LastName = parts[2]
	}
}

// GatherEmails tries to get the emails from the description
func (r *ServiceRepository) GatherEmails(service *Service) {
	re := regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
	emails := re.FindAllString(service.Description, -1)
	for _, email := range emails {
		service.Emails = append(service.Emails, Email{Email: email})
	}
}

// GatherPhones tries to get the phones from the description
func (r *ServiceRepository) GatherPhones(service *Service) {
	const minLen = 4
	re := regexp.MustCompile(`(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?`)
	phones := re.FindAllString(service.Description, -1)
	for _, phone := range phones {
		phone = strings.TrimSpace(phone)
		if len(phone) > minLen {
			service.Phones = append(service.Phones, Phone{Phone: phone})
		}
	}
}

// NewServiceRepository return a new ServiceRepository
func NewServiceRepository(database Database) *ServiceRepository {
	return &ServiceRepository{db: database}
}
