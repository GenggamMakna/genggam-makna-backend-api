package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email       string
	Password    string
	FirstName   string
	LastName    string
	GoogleToken string
}
