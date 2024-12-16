package service

import (
	"fmt"
	"log/slog"
	"net/smtp"
)

// Config для настройки SMTP
type EmailConfig struct {
	SMTPHost     string // SMTP-сервер (например, "smtp.gmail.com")
	SMTPPort     string // Порт SMTP (например, "587")
	SMTPUsername string // Логин (email-адрес отправителя)
	SMTPPassword string // Пароль от почты
}

// SendEmail отправляет письмо на указанный email-адрес
func SendEmail(logger *slog.Logger, config EmailConfig, to, subject, body string) error {

	logger.Debug("Функция отправки сообщения SendEmail")
	// Создаем сообщение
	message := fmt.Sprintf("From: %s\r\n", config.SMTPUsername) +
		fmt.Sprintf("To: %s\r\n", to) +
		fmt.Sprintf("Subject: %s\r\n\r\n", subject) +
		body

	logger.Debug("Сформированное сообщение: ", message)

	logger.Debug("Настраиваем аутентификацию")
	// Настраиваем аутентификацию
	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPHost)

	logger.Debug("Отправляем письмо")
	// Отправляем письмо
	err := smtp.SendMail(
		config.SMTPHost+":"+config.SMTPPort, // Адрес SMTP-сервера
		auth,                                // Аутентификация
		config.SMTPUsername,                 // Отправитель
		[]string{to},                        // Получатель
		[]byte(message),                     // Тело письма
	)
	if err != nil {
		logger.Error("Ошибка при отправке сообщения")
		logger.Error(err.Error())
		return fmt.Errorf("failed to send email: %w", err)
	}

	logger.Debug("Функция отправки сообщения (SendEmail) отработала успешно!")
	return nil
}
