package services

import (
	"context"
	"encoding/json"
	"fmt"
	"genggam-makna-api/config"
	"genggam-makna-api/dto"
	"time"

	"github.com/go-redis/redis/v8"
)

func (s *compServices) CachePredict(image []byte, data dto.MLResponse, model_type dto.ModelType) error {
	rdb := config.ConnectRedis()
	ctx := context.Background()

	jsonValue, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling value: " + err.Error())
	}

	err = rdb.Set(ctx, string(model_type)+string(image), jsonValue, time.Hour*2).Err()
	if err != nil {
		fmt.Println("Error setting key in Redis: " + err.Error())
	}

	return nil
}

func (s *compServices) GetPredictCache(image []byte, model_type dto.ModelType) (*dto.MLResponse, error) {
	rdb := config.ConnectRedis()
	ctx := context.Background()

	result, err := rdb.Get(ctx, string(model_type)+string(image)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		fmt.Printf("Error getting value from Redis: %v\n", err)
		return nil, err
	}

	var response dto.MLResponse
	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		fmt.Printf("Error unmarshaling value: %v\n", err)
		return nil, err
	}

	return &response, nil
}
