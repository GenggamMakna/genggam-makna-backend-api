package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"cloud.google.com/go/storage"
)

func (s *compServices) StoreImage(image []byte) (string, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to generate random name: %w", err)
	}
	randomName := hex.EncodeToString(randomBytes) + ".png"

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create storage client: %w", err)
	}
	defer client.Close()

	bucket := client.Bucket("xanny-bucket")

	writer := bucket.Object(randomName).NewWriter(ctx)
	defer writer.Close()

	if _, err := writer.Write(image); err != nil {
		return "", fmt.Errorf("failed to write image to bucket: %w", err)
	}

	return randomName, nil
}
