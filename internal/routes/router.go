package routes

import (
	"os"
	"pulsy/internal/handlers"
	"pulsy/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Set the Gin to release mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()

	// Health Check
	router.GET("/health", handlers.HealthCheckHandler)

	protected := router.Group("/api", middleware.VerifyTokenOfFirebase)

	protected.POST("/upload", handlers.UploadFileHandler)
	protected.GET("/download/:uuid", handlers.DownloadFileHandler)

	// Auth Routes
	router.GET("/user/:uid", handlers.GetUser)

	router.POST("/token", handlers.GenerateAccessTokenByUserID)

	// Websocket routes
	protected.GET("/ws", handlers.Websocket)

	return router
}
