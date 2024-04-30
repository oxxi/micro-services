package services

import (
	"context"
	"sync"

	"github.com/oxxi/cactus-tech/internal/dto"
	"github.com/oxxi/cactus-tech/internal/entities"
	"github.com/oxxi/cactus-tech/internal/repository"
)

type IMetricService interface {
	GetAllMetric(ctx context.Context) ([]dto.Metric, error)
}

type metricService struct {
	repo repository.IMetricsRepository
}

// GetAllMetric implements IMetricService.
func (m *metricService) GetAllMetric(ctx context.Context) ([]dto.Metric, error) {

	entity, err := m.repo.GetAllMetrics(ctx)
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, len(entity))

	for _, entityMetric := range entity {
		metricDto := dto.Metric{
			Name:        entityMetric.Name,
			Measurement: toMeasurementResponse(entityMetric.Measurement)}
		metrics = append(metrics, metricDto)
	}

	return metrics, nil
}

func toMeasurementResponse(measurements []entities.Measurement) []dto.MeasurementResponse {
	responses := make([]dto.MeasurementResponse, 0, len(measurements))
	for _, measurement := range measurements {
		response := dto.MeasurementResponse{
			Value:     uint(measurement.Value),
			Timestamp: measurement.Timestamp,
		}
		responses = append(responses, response)
	}
	return responses
}

var once sync.Once
var instance *metricService

func NewMetricService(r repository.IMetricsRepository) IMetricService {
	once.Do(func() {
		instance = &metricService{
			repo: r,
		}
	})
	return instance
}
