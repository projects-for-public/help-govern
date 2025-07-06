package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/projects-for-public/help-govern/internal/config"
	"github.com/projects-for-public/help-govern/internal/handlers"
	"github.com/projects-for-public/help-govern/internal/services"
	"github.com/projects-for-public/help-govern/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		utils.Fatal("Failed to load config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		utils.Fatal("Failed to connect to database: %v", err)
	}

	reportService := services.NewReportService(db)
	reportHandler := handlers.NewReportHandler(reportService)

	h := &handlers.Handlers{
		Report: reportHandler,
		// Add other handlers here as needed
	}

	r := gin.Default()

	utils.Info("Registering routes and starting server...")
	handlers.RegisterRoutes(r, h)

	if err := r.Run(); err != nil {
		utils.Error("Server failed to start: %v", err)
		os.Exit(1)
	}
}
