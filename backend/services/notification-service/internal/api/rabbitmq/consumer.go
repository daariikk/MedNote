package rabbitmq

import (
	"encoding/json"
	"github.com/daariikk/MedNote/services/notification-service/internal/config"
	"github.com/daariikk/MedNote/services/notification-service/internal/domain"
	"github.com/daariikk/MedNote/services/notification-service/internal/repository/postgres"
	"github.com/daariikk/MedNote/services/notification-service/internal/service"
	"log/slog"
)

// StartConsumer запускает прослушивание очереди RabbitMQ
func StartConsumer(logger *slog.Logger, db *postgres.Storage, rabbit *RabbitMQ, cfg *config.Config) error {
	ch, err := rabbit.Channel.Consume(
		"reset-password", // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range ch {
			logger.Info("Received a message", slog.String("body", string(d.Body)))

			// Обработка сообщения
			var user domain.User
			if err := json.Unmarshal(d.Body, &user); err != nil {
				logger.Error("Failed to unmarshal message", slog.String("error", err.Error()))
				continue
			}

			// Отправка уведомления
			if err := service.SendNotification(logger, db, &user, cfg); err != nil {
				logger.Error("Failed to send notification", slog.String("error", err.Error()))
			}
		}
	}()

	logger.Info("Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}
