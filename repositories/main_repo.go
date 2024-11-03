package repositories

import (
	"genggam-makna-api/config"

	"gorm.io/gorm"
)

type CompRepository interface{}

type compRepository struct {
	DB *gorm.DB
}

func NewComponentRepository(DB *gorm.DB) *compRepository {
	db := config.InitDB()

	return &compRepository{
		DB: db,
	}
}
