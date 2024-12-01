package app

import (
	"context"
	"database/sql"
	"github.com/daariikk/MedNote/services/patient-service/internal/config"
	"github.com/daariikk/MedNote/services/patient-service/internal/lib/logger/handlers/slogpretty"
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
	db     *sql.DB
}

func New(config *config.Config) *App {
	logger := setupLogger(config.Env)

	return &App{
		logger: logger,
		config: config,
		server: &http.Server{
			Addr:         config.Address,
			ReadTimeout:  config.Timeout,
			WriteTimeout: config.Timeout,
			IdleTimeout:  config.IdleTimeout,
		},
	}
}

func (a *App) Run() {

	go func() {
		a.logger.Info("Starting server", slog.String("addr", a.config.Address))
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Error("Server failed to start", slog.String("error", err.Error()))
		}
	}()

	//Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	a.logger.Info("Shutting down server...")

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
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
