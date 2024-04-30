package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oxxi/cactus-tech/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMetricService struct {
	mock.Mock
}

func (m *MockMetricService) GetAllMetric(ctx context.Context) ([]dto.Metric, error) {
	args := m.Called(ctx)
	return args.Get(0).([]dto.Metric), args.Error(1)
}

func TestGetMetrics(t *testing.T) {
	mockService := new(MockMetricService)
	handler := NewMetricHandler(mockService)

	// Test de éxito
	t.Run("success", func(t *testing.T) {
		measu := make([]dto.MeasurementResponse, 0, 1)
		m := dto.MeasurementResponse{Value: 1, Timestamp: time.Now()}
		measu = append(measu, m)
		expectedMetrics := []dto.Metric{{Name: "metric1", Measurement: measu}}
		mockService.On("GetAllMetric", mock.Anything).Return(expectedMetrics, nil)

		req, _ := http.NewRequest("GET", "/metrics", nil)
		rr := httptest.NewRecorder()

		httpHandler := http.HandlerFunc(handler.GetMetrics)
		httpHandler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		expectedJSON, _ := json.Marshal(expectedMetrics)
		assert.JSONEq(t, string(expectedJSON), rr.Body.String())
		mockService.AssertExpectations(t)
	})
}

func TestGetMetrics_InternalError(t *testing.T) {
	mockService := new(MockMetricService)
	handler := NewMetricHandler(mockService)
	t.Run("service failure", func(t *testing.T) {
		mockService.On("GetAllMetric", mock.Anything).Return([]dto.Metric{}, errors.New("failed to retrieve metrics"))

		badRequest, _ := http.NewRequest("GET", "/metrics", nil)
		responseRequest := httptest.NewRecorder()

		httpHandler := http.HandlerFunc(handler.GetMetrics)
		httpHandler.ServeHTTP(responseRequest, badRequest)

		assert.Equal(t, http.StatusInternalServerError, responseRequest.Code)
		assert.Contains(t, responseRequest.Body.String(), "Failed to retrieve metrics")
		mockService.AssertExpectations(t)
	})
}
