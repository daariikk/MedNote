package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/daariikk/MedNote/services/api-gateway/internal/api/response"
	"github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/helper"
	"github.com/daariikk/MedNote/services/api-gateway/internal/config"
	"github.com/daariikk/MedNote/services/api-gateway/internal/domain"
	"log/slog"
	"net/http"
)

type LoginWrapper interface {
	GetPassword(string) (int64, string, error)
}

func LoginHandler(logger *slog.Logger, auth LoginWrapper, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("LoginHandler starting...")

		request := domain.User{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		logger.Debug("Body успешно распарсен")

		logger.Debug("Пытаемся получить пароль по указанному email")
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
		if decodedPassword != request.Password {
			logger.Info("Пароли не совпадают")
			logger.Error(err.Error())
			response.SendFailureResponse(w, fmt.Sprintf("Failed to auth user: %v", err), http.StatusUnauthorized)
			return
		}
		logger.Debug("Пароль введен успешно")

		logger.Debug("Пытаемся снегерировать токен")
		token, err := generateToken(cfg, patientId)
		if err != nil {
			logger.Error("Произошла ошибка внутри функции generateToken")
			logger.Error("Failed to generate token", slog.String("error", err.Error()))
			response.SendFailureResponse(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}
		logger.Debug("Токен успешно сгенерирован")

		logger.Debug("Формируем ответ")
		res := map[string]interface{}{
			"patient_id": patientId,
			"token":      token,
		}
		logger.Debug("Сформированный ответ", res)

		logger.Info("LoginHandler works successful")
		response.SendSuccessResponse(w, res, http.StatusOK)
	}
}
