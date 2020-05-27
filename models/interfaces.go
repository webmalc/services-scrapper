package models

import "github.com/jinzhu/gorm"

// AutoMigrater auto migrate the DB
type AutoMigrater interface {
	AutoMigrate(values ...interface{}) *gorm.DB
}

// NameProcessor processes names
type NameProcessor interface {
	ProcessName(service *Service)
}

// PhonesGatherer gather phones
type PhonesGatherer interface {
	GatherPhones(service *Service)
}

// EmailsGatherer gather emails
type EmailsGatherer interface {
	GatherEmails(service *Service)
}

// ServiceProcessor interface
type ServiceProcessor interface {
	NameProcessor
	EmailsGatherer
	PhonesGatherer
}

// Creator interface
type Creator interface {
	Create(value interface{}) *gorm.DB
}

// Finder interface
type Finder interface {
	Model(value interface{}) *gorm.DB
}

// Database interface
type Database interface {
	Creator
	Finder
}
