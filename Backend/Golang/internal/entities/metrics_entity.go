package entities

type Metric struct {
	ID          uint          `gorm:"primarykey"`
	Name        string        `gorm:"type:varchar(255);not null"`
	Description string        `gorm:"type:text;not null"`
	Measurement []Measurement `gorm:"foreignKey:MetricID"`
}
