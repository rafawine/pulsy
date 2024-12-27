package handlers

import (
	"net/http"

	"pulsy/internal/responses"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, responses.Success("Pulsy API is running", nil))
}
