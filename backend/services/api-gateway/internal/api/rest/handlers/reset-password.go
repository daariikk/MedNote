package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/daariikk/MedNote/services/api-gateway/internal/api/response"
	"github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/helper"
	"github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/rabbitmq"
	"github.com/daariikk/MedNote/services/api-gateway/internal/config"
	"log/slog"
	"math/big"
	"net/http"
	"os"
)

type UpdatePasswordWrapper interface {
	UpdatePassword(email string, password string) error
}

type EmailRequest struct {
	Email string `json:"email"`
}

type UserRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

// NewPassword генерирует случайный пароль длиной 10 символов
func NewPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const passwordLength = 10

	password := make([]byte, passwordLength)

	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		password[i] = charset[randomIndex.Int64()]
	}
	return string(password)
}

// ResetHandler обрабатывает запрос на сброс пароля
func ResetHandler(logger *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("ResetHandler starting...")
		var request EmailRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			logger.Error("Failed to decode request body", slog.String("error", err.Error()))
			return
		}

		logger.Debug("Body успешно распарсен", slog.String("email", request.Email))

		newPass := NewPassword()
		if newPass == "" {
			http.Error(w, "Failed to generate password", http.StatusInternalServerError)
			logger.Error("Failed to generate password")
			return
		}

		logger.Debug("New password generated", slog.String("password", newPass))
		// Подключаемся к RabbitMQ
		rabbit, err := rabbitmq.Connect(cfg.RabbitMQ.URL)
		if err != nil {
			http.Error(w, "Failed to connect to RabbitMQ", http.StatusInternalServerError)
			logger.Error("Failed to connect to RabbitMQ", slog.String("error", err.Error()))
			return
		}
		defer rabbit.Close()

		// Подготовка сообщения для отправки в RabbitMQ
		message := map[string]string{
			"email":    request.Email,
			"password": newPass,
		}
		messageBytes, err := json.Marshal(message)
		if err != nil {
			http.Error(w, "Failed to prepare message", http.StatusInternalServerError)
			logger.Error("Failed to marshal message", err, slog.String("error", err.Error()))
			return
		}

		// Отправляем сообщение в очередь reset-password
		if err := rabbitmq.Publish(rabbit, "reset-password", messageBytes); err != nil {
			http.Error(w, "Failed to send message to RabbitMQ", http.StatusInternalServerError)
			logger.Error("Failed to publish message to RabbitMQ", err, slog.String("error", err.Error()))
			return
		}

		logger.Info("Message sent to RabbitMQ", slog.String("email", request.Email))

		// Возвращаем успешный ответ
		response.SendSuccessResponse(w, "Password reset request processed", http.StatusCreated)

	}
}

func RestorationHandler(logger *slog.Logger, auth LoginWrapper, reset UpdatePasswordWrapper, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("RestorationHandler starting...")

		var request UserRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			logger.Error("Failed to decode request body", slog.String("error", err.Error()))
			return
		}

		patientId, encodedPassword, err := auth.GetPassword(request.Email)
		if err != nil {
			logger.Error("Произошла ошибка внутри функции GetPassword")
			logger.Error(err.Error())
			response.SendFailureResponse(w, fmt.Sprintf("Failed to auth user: %v", err), http.StatusInternalServerError)
			return
		}
		logger.Debug("GetPassword отработала успешно")
		logger.Debug("patientId и encodedPassword", slog.Int64("patientId", patientId), slog.String("encodedPassword", encodedPassword))

		logger.Debug("Пытаемся расшифровать пароль")
		decodedPassword, err := helper.DecodePassword(encodedPassword)
		if err != nil {
			logger.Error("Произошла ошибка внутри функции DecodePassword")
			logger.Error(err.Error())
			response.SendFailureResponse(w, fmt.Sprintf("Failed to auth user: %v", err), http.StatusInternalServerError)
			return
		}
		logger.Debug("Пароль успешно расшифрован")

		logger.Debug("Пытаемся проверить совпадают ли пароли")
		logger.Debug("Расшифрованный пароль: ", slog.String("decodedPassword", decodedPassword))
		logger.Debug("Переданный пароль: ", slog.String("request.Password", request.Password))

		if decodedPassword != request.Password {
			logger.Info("Пароли не совпадают")
			logger.Error(err.Error())
			response.SendFailureResponse(w, fmt.Sprintf("Failed to auth user: %v", err), http.StatusUnauthorized)
			return
		}
		logger.Debug("Пароль введен успешно")

		encodedPassword = helper.EncodePassword(request.NewPassword)

		if err := reset.UpdatePassword(request.Email, encodedPassword); err != nil {
			logger.Error("Ошибка обновления пароля")
			response.SendFailureResponse(w, fmt.Sprintf("Failed to reset password: %v", err), http.StatusInternalServerError)
			return
		}
		logger.Debug("RestorationHandler works successful")
		response.SendSuccessResponse(w, "Password reset request processed", http.StatusCreated)
	}
}

func GetRestorationHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("GetRestorationHandler starting...")

		// Указываем абсолютный путь к файлу restoration.html
		restorationPath := "C:\\Users\\Daria\\Projects All\\Go Projects\\MedNote\\frontend\\templates\\auth\\restoration.html"
		logger.Debug("restorationPath:", "restorationPath", restorationPath)

		// Проверяем, существует ли файл
		if _, err := os.Stat(restorationPath); os.IsNotExist(err) {
			logger.Error("File not found", "path", restorationPath, "error", err)
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		// Отдаем файл
		http.ServeFile(w, r, restorationPath)
		logger.Debug("GetRestorationHandler works successful")
	}
}
