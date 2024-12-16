package service

import (
	"encoding/base64"
	"fmt"
	"github.com/daariikk/MedNote/services/notification-service/internal/config"
	"github.com/daariikk/MedNote/services/notification-service/internal/domain"
	"github.com/daariikk/MedNote/services/notification-service/internal/repository/postgres"
	"log/slog"
)

func EncodePassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}

// SendNotification отправляет уведомление пользователю
func SendNotification(logger *slog.Logger, db *postgres.Storage, user *domain.User, cfg *config.Config) error {
	logger.Info("Sending notification", slog.String("email", user.Email))

	encodedPass := EncodePassword(user.Password)

	logger.Debug("Устанавливаем временный пароль в Базе данных")
	// Устанавливаем временный пароль в Базе данных
	if err := db.UpdatePassword(user.Email, encodedPass); err != nil {
		return err
	}

	logger.Debug("Временный пароль успешно установлен")

	emailConfig := EmailConfig{
		SMTPHost:     cfg.SMTP.Host,
		SMTPPort:     cfg.SMTP.Port,
		SMTPUsername: cfg.SMTP.Username,
		SMTPPassword: cfg.SMTP.Password,
	}
	logger.Debug("Конфигурационные настройки для отправки сообщения: ", slog.Any("config", emailConfig))

	subject := "MedNote. Восстановление пароля"
	body := fmt.Sprintf("Здравствуйте! \n\nДля восстановления пароля перейдите по ссылке: http://localhost:8082/api/v1/restoration\nВаш временный пароль: %s", user.Password)
	if err := SendEmail(logger, emailConfig, user.Email, subject, body); err != nil {
		logger.Error("Failed to send email", slog.String("error", err.Error()))
		return err
	}

	logger.Info("Notification sent successfully", slog.String("email", user.Email))
	return nil
}
