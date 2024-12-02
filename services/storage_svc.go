package services

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"
)

func (s *compServices) StoreImage(image []byte) (string, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	randomName := hex.EncodeToString(randomBytes) + ".png"

	publicFolderPath := "./public"
	if err := os.MkdirAll(publicFolderPath, os.ModePerm); err != nil {
		return "", err
	}

	filePath := filepath.Join(publicFolderPath, randomName)

	if err := os.WriteFile(filePath, image, 0644); err != nil {
		return "", err
	}

	return randomName, nil
}
