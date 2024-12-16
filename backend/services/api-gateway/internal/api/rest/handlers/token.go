package handlers

import (
	"fmt"
	"github.com/daariikk/MedNote/services/api-gateway/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func generateToken(cfg *config.Config, patientID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"patient_id": patientID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte(cfg.JWT.SecretKey)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(cfg *config.Config, tokenString string) (*jwt.Token, error) {
	secretKey := []byte(cfg.JWT.SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
