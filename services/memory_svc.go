package services

import (
	"context"
	"encoding/json"
	"fmt"
	"genggam-makna-api/config"
	"genggam-makna-api/dto"
)

func (s *compServices) CachePredict(image []byte, data dto.MLResponse) error {
	rdb := config.ConnectRedis()
	ctx := context.Background()

	jsonValue, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling value: " + err.Error())
	}

	err = rdb.Set(ctx, string(image), jsonValue, 0).Err()
	if err != nil {
		fmt.Println("Error setting key in Redis: " + err.Error())
	}

	return nil
}