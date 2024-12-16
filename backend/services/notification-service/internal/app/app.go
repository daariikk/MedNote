package app

import (
	"context"
	"fmt"
	"github.com/daariikk/MedNote/services/notification-service/internal/api/rabbitmq"
	"github.com/daariikk/MedNote/services/notification-service/internal/config"
	"github.com/daariikk/MedNote/services/notification-service/internal/lib/logger/handlers/slogpretty"
	"github.com/daariikk/MedNote/services/notification-service/internal/repository/postgres"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	logger *slog.Logger
	config *config.Config
	server *http.Server
	db     *postgres.Storage
	rabbit *rabbitmq.RabbitMQ
}

type DBParams struct {
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Port     string `env:"DB_PORT"`
}

func New(config *config.Config) *App {
	logger := setupLogger(config.Env)

	dbUrlConnection := CreateDBConnectionUrl(*config)
	logger.Debug("DB url connection", slog.String("url", dbUrlConnection))
	ctx := context.Background()

	storage, err := postgres.New(ctx, logger, dbUrlConnection)
	if err != nil {
		log.Fatalf("db create error: %s", err)
	}

	rabbit, err := rabbitmq.Connect(config.RabbitMQ.URL)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %s", err)
	}

	return &App{
		logger: logger,
		config: config,
		server: &http.Server{
			Addr:         config.Address,
			ReadTimeout:  config.Timeout,
			WriteTimeout: config.Timeout,
			IdleTimeout:  config.IdleTimeout,
		},
		db:     storage,
		rabbit: rabbit,
	}
}

func (a *App) Run() {
	defer a.db.Close()
	defer a.rabbit.Close()

	errChan := make(chan error, 1)
	go func() {
		a.logger.Info("Starting server", slog.String("addr", a.config.Address))
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	go func() {
		a.logger.Info("Starting RabbitMQ consumer")
		if err := rabbitmq.StartConsumer(a.logger, a.db, a.rabbit, a.config); err != nil {
			errChan <- err
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		a.logger.Error("Server failed to start", slog.String("error", err.Error()))
		return
	case <-quit:
		a.logger.Info("Shutting down server...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error("Server forced to shutdown", slog.String("error", err.Error()))
	}

	a.logger.Info("Server exiting")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func CreateDBConnectionUrl(cfg config.Config) string {
	paramsDB := DBParams{}
	err := cleanenv.ReadEnv(&paramsDB)
	if err != nil {
		log.Fatalf("Error reading DB connection params: %s", err)
	}

	return fmt.Sprintf(cfg.DatabaseBaseUrl, paramsDB.Username, paramsDB.Password, paramsDB.Port)
}
