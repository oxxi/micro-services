package handler

import (
	"encoding/json"
	"net/http"

	"github.com/oxxi/cactus-tech/internal/services"
)

type IMetricHandler interface {
	GetMetrics(w http.ResponseWriter, r *http.Request)
}

type metricHandler struct {
	service services.IMetricService
}

// GetMetrics implements IMetricHandler.
func (m *metricHandler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := m.service.GetAllMetric(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to retrieve metrics", http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to encode metrics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func NewMetricHandler(s services.IMetricService) IMetricHandler {
	return &metricHandler{
		service: s,
	}
}
