package dto

import "time"

type MetricResponse struct {
	Cpu    Metric `json:"cpu"`
	Memory Metric `json:"memory"`
}

type Metric struct {
	Name        string                `json:"name"`
	Measurement []MeasurementResponse `json:"measurements"`
}

type MeasurementResponse struct {
	Value     uint      `json:"value"`
	Timestamp time.Time `json:"date"`
}
