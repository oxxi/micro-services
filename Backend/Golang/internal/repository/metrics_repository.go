package repository

import (
	"context"
	"fmt"

	"github.com/oxxi/cactus-tech/internal/entities"
	"gorm.io/gorm"
)

type IMetricsRepository interface {
	GetAllMetrics(ctx context.Context) ([]entities.Metric, error)
}

type metricRepository struct {
	Db *gorm.DB
}

// GetAllMetrics implements IMetricsRepository.
func (m *metricRepository) GetAllMetrics(ctx context.Context) ([]entities.Metric, error) {
	var metrics []entities.Metric
	result := m.Db.Model(&entities.Metric{}).WithContext(ctx).Preload("Measurement").Find(&metrics)
	fmt.Printf("result: %+v\n", result)
	if result.Error != nil {
		return nil, result.Error
	}
	return metrics, nil
}

func NewMetricRepository(db *gorm.DB) IMetricsRepository {
	return &metricRepository{
		Db: db,
	}
}
