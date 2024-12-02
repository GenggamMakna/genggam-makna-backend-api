package helpers

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/nfnt/resize"
)


func DownScaleImage(imageData []byte) ([]byte, error) {
	format, err := detectImageFormat(imageData)
	if err != nil {
		return nil, fmt.Errorf("unsupported image format: %w", err)
	}

	var img image.Image
	switch format {
	case "png":
		img, err = png.Decode(bytes.NewReader(imageData))
	case "jpeg":
		img, err = jpeg.Decode(bytes.NewReader(imageData))
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
	if err != nil {
		return nil, fmt.Errorf("error decoding image: %w", err)
	}

	resizedImg := resize.Resize(244, 244, img, resize.Lanczos3)

	var buffer bytes.Buffer
	err = png.Encode(&buffer, resizedImg)
	if err != nil {
		return nil, fmt.Errorf("error encoding image: %w", err)
	}

	return buffer.Bytes(), nil
}

func detectImageFormat(imageData []byte) (string, error) {
	_, format, err := image.DecodeConfig(bytes.NewReader(imageData))
	if err != nil {
		return "", err
	}
	return format, nil
}
