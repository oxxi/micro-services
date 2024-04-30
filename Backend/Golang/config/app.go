package config

import (
	"fmt"
	"log"
	"os"

	"github.com/oxxi/cactus-tech/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func Connect() {

	//"host=127.0.0.1 user=root password=123456789 dbname=db_metrics port=5432 sslmode=disable"
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, db_name, port)
	log.Println(dsn)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	d.AutoMigrate(&entities.Metric{}, &entities.Measurement{})
	db = d
	go setInitValues()
}

func GetDB() *gorm.DB {
	return db
}

func setInitValues() {
	//Metric
	var metrics []entities.Metric
	result := db.Find(&metrics)
	if result.RowsAffected == 0 {
		memory := entities.Metric{ID: 1, Name: "Memory", Description: "Memory usage"}
		cpu := entities.Metric{ID: 2, Name: "CPU", Description: "CPU usage"}
		db.Create(&cpu)
		db.Create(&memory)

	}

}
