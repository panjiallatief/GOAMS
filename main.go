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
	db.AutoMigrate(&models.Asset{})

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

	// Asset endpoints
	assetRepo := repository.NewAssetRepository(db)
	assetService := service.NewAssetService(assetRepo)
	handlers.NewAssetHandler(router, assetService)

	// TODO: add other handlers

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
