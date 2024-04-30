package pkg

import (
	"net/http"

	"github.com/oxxi/cactus-tech/internal/handler"
	"github.com/oxxi/cactus-tech/internal/repository"
	"github.com/oxxi/cactus-tech/internal/services"
	"gorm.io/gorm"
)

var RegisterRouter = func(router *http.ServeMux, db *gorm.DB) {

	repo := repository.NewMetricRepository(db)
	services := services.NewMetricService(repo)
	handler := handler.NewMetricHandler(services)

	router.HandleFunc("GET /", handler.GetMetrics)

}
