package repositories

import (
	"errors"
	"genggam-makna-api/dto"
	"genggam-makna-api/models"
	"strings"
)

func (r *compRepository) RegisterUserCredential(data dto.User) (string, error) {
	user_data := models.Users{
		Email:     data.Email,
		Password:  data.Password,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	result := r.DB.Create(&user_data)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return "", errors.New("409")
		}
		return "", result.Error
	}

	return user_data.ID.String(), nil
}

func (r *compRepository) LoginUserCredentials(email string) (*models.Users, error) {
	var user_data models.Users
	result := r.DB.Where("email = ?", email).First(&user_data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user_data, nil
}
