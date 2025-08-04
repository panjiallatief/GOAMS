package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/panjiallatief/GOAMS/config"
	"github.com/panjiallatief/GOAMS/handlers"
	"github.com/panjiallatief/GOAMS/models"
	"github.com/panjiallatief/GOAMS/repository"
	"github.com/panjiallatief/GOAMS/service"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate database
	db.AutoMigrate(&models.Asset{}, &models.Log{})

	router := gin.Default()

	// Serve dashboard static files
	router.Static("/src", "./src")
	router.GET("/", func(c *gin.Context) {
		c.File("./src/index.html")
	})

	// Health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Initialize repositories
	assetRepo := repository.NewAssetRepository(db)
	logRepo := repository.NewLogRepository(db)
	dropdownRepo := repository.NewDropdownRepository(db)

	// Initialize services
	assetService := service.NewAssetService(assetRepo)
	logService := service.NewLogService(logRepo)
	dropdownService := service.NewDropdownService(dropdownRepo)

	// Initialize handlers
	handlers.NewAssetHandler(router, assetService, logService)
	handlers.NewDropdownHandler(router, dropdownService)

	// TODO: add other handlers

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
