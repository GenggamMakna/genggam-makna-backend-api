package repositories

import (
	"genggam-makna-api/config"
	"genggam-makna-api/dto"
	"genggam-makna-api/models"

	"gorm.io/gorm"
)

type CompRepository interface {
	RegisterUserCredential(data dto.User) (string, error)
	LoginUserCredentials(email string) (*models.Users, error)
	LoginUserGoogle(data dto.User) (string, error)
}

type compRepository struct {
	DB *gorm.DB
}

func NewComponentRepository(DB *gorm.DB) *compRepository {
	db := config.InitDB()

	return &compRepository{
		DB: db,
	}
}
