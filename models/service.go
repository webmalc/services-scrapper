package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
)

// Phone is the email struct
type Phone struct {
	gorm.Model
	Phone     string `gorm:"type:varchar(100);not null;index:phone"`
	ServiceID uint
}

// Email is the email struct
type Email struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);not null;index:email" valid:"email"`
	ServiceID uint
}

// Image is the image struct
type Image struct {
	gorm.Model
	Src       string `gorm:"size:255;not null;index:src" valid:"url"`
	ServiceID uint
}

// Link is the link struct
type Link struct {
	gorm.Model
	URL       string `gorm:"size:255;not null;" valid:"url"`
	ServiceID uint
}

// Service is the service struct
type Service struct {
	gorm.Model
	URL         string           `gorm:"size:255;not null;unique_index;" valid:"url"`
	Source      string           `gorm:"size:255;not null;index:source;" valid:"required"`
	Title       string           `gorm:"size:255;not null;" valid:"required"`
	Description string           `gorm:"type:text;index:description;"`
	FullName    string           `gorm:"type:varchar(255);"`
	FirstName   string           `gorm:"type:varchar(255);"`
	MiddleName  string           `gorm:"type:varchar(255);"`
	LastName    string           `gorm:"type:varchar(255);"`
	Address     string           `gorm:"index:addr"`
	Phones      []Phone          `gorm:"foreignkey:ServiceID"`
	Emails      []Email          `gorm:"foreignkey:ServiceID"`
	Images      []Image          `gorm:"foreignkey:ServiceID"`
	Links       []Link           `gorm:"foreignkey:ServiceID"`
	Processor   ServiceProcessor `gorm:"-"`
}

// Validate validates struct
func (s *Service) Validate(db *gorm.DB) {
	config := NewConfig()
	if !govalidator.IsIn(s.Source, config.scrappers...) {
		_ = db.AddError(
			validations.NewError(s, "Source", "the source is inccorrect"),
		)
	}
}

// BeforeSave runs before saving the object
func (s *Service) BeforeSave() (err error) {
	s.Processor.ProcessName(s)
	s.Processor.GatherEmails(s)
	s.Processor.GatherPhones(s)
	return nil
}

// Migrate migrates the DB
func Migrate(migrater AutoMigrater) {
	migrater.AutoMigrate(&Service{}, &Phone{}, &Email{}, &Image{}, &Link{})
}
