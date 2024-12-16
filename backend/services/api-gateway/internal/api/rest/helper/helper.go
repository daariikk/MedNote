package helper

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func EncodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}

func DecodePassword(encodedPassword string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedPassword)
	if err != nil {
		return "", fmt.Errorf("failed to decode password: %w", err)
	}
	return string(decodedBytes), nil
}

func CopyHeaders(dst, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dst.Add(key, value)
		}
	}
}
