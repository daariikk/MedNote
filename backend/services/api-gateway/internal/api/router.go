package api

import (
	"github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/handlers"
	patient_service "github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/handlers/patient-service"
	record_service "github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/handlers/record-service"
	reminder_service "github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/handlers/reminder-service"
	"github.com/daariikk/MedNote/services/api-gateway/internal/config"
	"github.com/daariikk/MedNote/services/api-gateway/internal/repository/postgres"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

const frontendDir = "C:\\Users\\Daria\\Projects All\\Go Projects\\MedNote\\frontend"

func NewRouter(cfg *config.Config, logger *slog.Logger, storage *postgres.Storage) *chi.Mux {
	router := chi.NewRouter()
	router.Use(handlers.CorsMiddleware)

	router.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", handlers.RegisterHandler(logger, storage))
		r.Post("/login", handlers.LoginHandler(logger, storage, cfg))
		r.Post("/reset-password", handlers.ResetHandler(logger, cfg))
		r.Post("/restoration", handlers.RestorationHandler(logger, storage, storage, cfg))
		r.Get("/restoration", handlers.GetRestorationHandler(logger))
	})

	router.Handle("/api/static/*", http.StripPrefix("/api/static/", http.FileServer(http.Dir(frontendDir))))

	router.Group(func(r chi.Router) {

		r.Use(handlers.AuthMiddleware(logger, cfg))

		r.Route("/api/v1/users/{userId}/patient", func(r chi.Router) {
			r.Get("/", patient_service.GetPatient(logger, cfg))
			r.Put("/", patient_service.UpdatePatient(logger, cfg))
		})

		r.Route("/api/v1/records", func(r chi.Router) {
			r.Post("/", record_service.NewRecord(logger, cfg))
			r.Get("/", record_service.GetRecords(logger, cfg))
			r.Delete("/", record_service.DeleteRecord(logger, cfg))
		})

		r.Route("/api/v1/reminders", func(r chi.Router) {
			r.Post("/", reminder_service.NewReminder(logger, cfg))
			r.Get("/", reminder_service.GetReminders(logger, cfg))
			r.Delete("/", reminder_service.DeleteReminder(logger, cfg))
		})
	})

	return router
}
