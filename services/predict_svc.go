package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"genggam-makna-api/dto"
	"genggam-makna-api/helpers"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func (s *compServices) ImagePredict(image_data []byte) (*dto.MLResponse, error) {
	predict_endpoint := os.Getenv("PREDICT_BASE_API_URL") + "/image"

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", "image.jpg")
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = part.Write(image_data)
	if err != nil {
		return nil, fmt.Errorf("failed to write file data: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	req, err := http.NewRequest("POST", predict_endpoint, &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var ml_response dto.MLResponse
	if err := json.Unmarshal(body, &ml_response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	go func() {
		compressed_file, err := helpers.DownScaleImage(image_data)
		if err != nil {
			fmt.Println(err)
		} else {

			result, err := s.StoreImage(compressed_file)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		}

		err = s.CachePredict(image_data, ml_response)
		if err != nil {
			fmt.Println(err)
		}
	}()

	return &ml_response, nil
}

func (s *compServices) VideoPredict(video_data []byte) (*dto.MLResponse, error) {
	predict_endpoint := os.Getenv("PREDICT_BASE_API_URL") + "/video"

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", "video.mp4")
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = part.Write(video_data)
	if err != nil {
		return nil, fmt.Errorf("failed to write video data: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	req, err := http.NewRequest("POST", predict_endpoint, &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var ml_response dto.MLResponse
	if err := json.Unmarshal(body, &ml_response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &ml_response, nil
}
