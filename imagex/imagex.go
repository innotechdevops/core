package imagex

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"
)

func ToBase64(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	var base64Encoding string
	mimeType := http.DetectContentType(b)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}
	base64Encoding += base64.StdEncoding.EncodeToString(b)
	return base64Encoding
}

func ToBytes(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		return []byte{}
	}
	return b
}

func ToImageBytes(path string, outputFormat string) []byte {
	inputBytes := ToBytes(path)
	img, _, err := image.Decode(bytes.NewReader(inputBytes))
	if err != nil {
		return []byte{}
	}

	var buf bytes.Buffer
	switch strings.ToLower(outputFormat) {
	case "png":
		err = png.Encode(&buf, img)
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
	default:
		return []byte{}
	}
	return buf.Bytes()
}

func ToImageScaleBytes(path string, maxWidth, maxHeight float64, outputFormat string) ([]byte, float64, float64) {
	imageBytes := ToImageBytes(path, outputFormat)
	if len(imageBytes) == 0 {
		return []byte{}, 0, 0
	}
	scaleX, scaleY, err := CalculateOptimalScale(imageBytes, maxWidth, maxHeight)
	if err != nil {
		return []byte{}, 0, 0
	}
	return imageBytes, scaleX, scaleY
}

func CalculateOptimalScale(imageBytes []byte, maxWidth, maxHeight float64) (float64, float64, error) {
	config, _, err := image.DecodeConfig(bytes.NewReader(imageBytes))
	if err != nil {
		return 1.0, 1.0, err
	}

	originalWidth := float64(config.Width)
	originalHeight := float64(config.Height)
	maxW := maxWidth
	maxH := maxHeight

	// Aspect ratio scaling
	scaleX := maxW / originalWidth
	scaleY := maxH / originalHeight

	// Choose the smaller scale to fit within both dimensions
	finalScale := scaleX
	if scaleY < scaleX {
		finalScale = scaleY
	}
	return finalScale, finalScale, nil
}
