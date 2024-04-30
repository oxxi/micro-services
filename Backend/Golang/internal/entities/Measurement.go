package entities

import (
	"time"
)

type Measurement struct {
	ID        uint      `gorm:"primarykey"`
	MetricID  uint      `gorm:"not null"`
	Value     uint      `gorm:"type:float;not null"`
	Timestamp time.Time `gorm:"type:timestamp;not null"`
}
