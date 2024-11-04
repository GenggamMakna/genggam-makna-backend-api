package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email       string `gorm:"unique;not null"`
	Password    string 
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	GoogleUID 	string 
}
