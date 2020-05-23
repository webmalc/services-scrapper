package models

import (
	"github.com/jinzhu/gorm"
)

// Service is the service struct
type Service struct {
	gorm.Model
	URL string `gorm:"size:255;not null;index:url;" valid:"url"`
}

// Migrate migrates the DB
func Migrate(migrater AutoMigrater) {
	migrater.AutoMigrate(&Service{})
}
