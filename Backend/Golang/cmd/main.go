package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/oxxi/cactus-tech/config"
	"github.com/oxxi/cactus-tech/pkg"
)

func main() {
	//
	fmt.Println(os.Getenv("APP_ENV"))
	if os.Getenv("APP_ENV") != "PRODUCTION" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	//set configuration
	config.Connect()
	db := config.GetDB()

	router := http.NewServeMux()
	go pkg.FakeData(db)
	pkg.RegisterRouter(router, db)

	corsRouter := pkg.EnableCors(router)

	if err := http.ListenAndServe(":8080", corsRouter); err != nil {
		panic(err)
	}
}
