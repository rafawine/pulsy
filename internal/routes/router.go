package routes

import (
	"os"
	"pulsy/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Set the Gin to release mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()

	// Health Check
	router.GET("/health", handlers.HealthCheckHandler)

	// File Routes
	router.POST("/upload", handlers.UploadFileHandler)
	router.GET("/download/:fileUUID", handlers.DownloadFileHandler)

	return router
}
