package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/panjiallatief/GOAMS/config"
	"github.com/panjiallatief/GOAMS/handlers"
	"github.com/panjiallatief/GOAMS/models"
	"github.com/panjiallatief/GOAMS/repository"
	"github.com/panjiallatief/GOAMS/service"
)

func main() {
	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate database
	db.AutoMigrate(&models.Asset{})

	router := gin.Default()

	// Serve frontend static files
	router.Static("/static", "./static")
	router.StaticFile("/", "./static/index.html")

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
