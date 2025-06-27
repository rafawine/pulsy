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

	// File routes V2
	fileRouter := router.Group("/file")
	fileRouter.POST("/", handlers.UploadFileHandler)
	fileRouter.GET("/:uuid", handlers.DownloadFileHandler)
	fileRouter.PUT("/:uuid", handlers.UpdateFileHandler)
	fileRouter.DELETE("/:uuid", handlers.DeleteFileHandler)

	return router
}
