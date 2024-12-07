package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/daariikk/MedNote/services/patient-service/internal/domain"
	"github.com/daariikk/MedNote/services/patient-service/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log/slog"
)

// MockGetPatienter - имитация GetPatienter для тестирования
type MockGetPatienter struct {
	mock.Mock
}

func (m *MockGetPatienter) GetPatient(userId int64) (domain.Patient, error) {
	args := m.Called(userId)
	return args.Get(0).(domain.Patient), args.Error(1)
}

func TestPatient(t *testing.T) {
	logger := slog.Default()

	t.Run("Invalid user ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/patient/invalidUserId", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := Patient(logger, nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("userId", "invalidUserId")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Contains(t, rr.Body.String(), "Invalid user ID")
	})

	t.Run("Patient not found", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/patient/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		mockPatienter := new(MockGetPatienter)
		mockPatienter.On("GetPatient", int64(1)).Return(domain.Patient{}, repository.ErrorNotFound)

		handler := Patient(logger, mockPatienter)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("userId", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Contains(t, rr.Body.String(), "Patient not found")
		mockPatienter.AssertExpectations(t)
	})

	t.Run("Internal server error", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/patient/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		mockPatienter := new(MockGetPatienter)
		mockPatienter.On("GetPatient", int64(1)).Return(domain.Patient{}, errors.New("some error"))

		handler := Patient(logger, mockPatienter)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("userId", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "Failed to get patient")
		mockPatienter.AssertExpectations(t)
	})

	t.Run("Successful getting patient", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/patient/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		mockPatienter := new(MockGetPatienter)
		expectedPatient := domain.Patient{
			Id:               1,
			FirstName:        "John",
			SecondName:       "Doe",
			Email:            "john.doe@example.com",
			Height:           180,
			Weight:           80,
			Gender:           "male",
			Password:         "password",
			RegistrationData: time.Now(),
		}
		mockPatienter.On("GetPatient", int64(1)).Return(expectedPatient, nil)

		handler := Patient(logger, mockPatienter)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("userId", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "John")
		mockPatienter.AssertExpectations(t)
	})
}
