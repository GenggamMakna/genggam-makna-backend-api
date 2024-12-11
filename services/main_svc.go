package services

import (
	"genggam-makna-api/dto"
	"genggam-makna-api/repositories"
)

type CompService interface {
	RegisterUserCredential(data dto.User) (*string, error)
	LoginUserCredentials(email string, password string) (*string, error)
	LoginUserGoogle(data dto.User) (*string, error)

	SIBIImagePredict(image_data []byte) (*dto.MLResponse, error)
	SIBIVideoPredict(video_data []byte) (*dto.MLResponse, error)
	
	BISINDOImagePredict(image_data []byte) (*dto.MLResponse, error)
	BISINDOVideoPredict(video_data []byte) (*dto.MLResponse, error)
	
	GetPredictCache(image []byte) (*dto.MLResponse, error)
}

type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}
