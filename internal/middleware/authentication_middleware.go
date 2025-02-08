package middleware

import (
	"log"
	"net/http"
	"pulsy/internal/responses"
	"pulsy/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyTokenOfFirebase(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	log.Println(authHeader)

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		message := "token not found"
		errData := map[string]interface{}{
			"cause": "access token is required",
		}

		c.IndentedJSON(http.StatusNotFound, responses.Error(message, errData))
		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	tokenData, err := services.VerifyToken(token)
	if err != nil {
		message := "invalid access token"
		errData := map[string]interface{}{
			"cause": err.Error(),
		}

		c.IndentedJSON(http.StatusNotFound, responses.Error(message, errData))
		c.Abort()
		return
	}

	c.Set("userID", tokenData.UID)

	c.Next()
}

func VerifyToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	log.Println(authHeader)

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		message := "token not found"
		errData := map[string]interface{}{
			"cause": "access token is required",
		}

		c.IndentedJSON(http.StatusNotFound, responses.Error(message, errData))
		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	tokenData, err := services.VerifyToken(token)
	if err != nil {
		message := "invalid access token"
		errData := map[string]interface{}{
			"cause": err.Error(),
		}

		c.IndentedJSON(http.StatusNotFound, responses.Error(message, errData))
		c.Abort()
		return
	}

	c.Set("userID", tokenData.UID)

	c.Next()
}
