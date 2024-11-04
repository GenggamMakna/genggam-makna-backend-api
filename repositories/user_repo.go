package repositories

import (
	"errors"
	"genggam-makna-api/dto"
	"genggam-makna-api/models"
	"strings"

	"gorm.io/gorm"
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

func (r *compRepository) LoginUserGoogle(data dto.User) (string, error) {
	var existingUser models.Users

	err := r.DB.Where("email = ?", data.Email).First(&existingUser).Error
	if err == nil {
		if existingUser.GoogleUID != "" {
			if existingUser.GoogleUID != data.GoogleUID {
				return "", errors.New("401")
			}
		} else if existingUser.Password != "" {
			existingUser.GoogleUID = data.GoogleUID
			if err := r.DB.Save(&existingUser).Error; err != nil {
				return "", err
			}
		}
		return existingUser.ID.String(), nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	userData := models.Users{
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		GoogleUID: data.GoogleUID,
	}

	result := r.DB.Create(&userData)
	if result.Error != nil {
		return "", result.Error
	}

	return userData.ID.String(), nil
}
