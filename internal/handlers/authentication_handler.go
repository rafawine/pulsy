package handlers

import (
	"net/http"
	"time"

	"pulsy/internal/requests"
	"pulsy/internal/responses"
	"pulsy/internal/services"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	uidUser := c.Param("uid")

	user, err := services.GetUser(uidUser)
	if err != nil {
		message := "user not found"
		errData := map[string]interface{}{
			"cause": err.Error(),
		}

		c.IndentedJSON(http.StatusNotFound, responses.Error(message, errData))
		return
	}

	message := "user found"
	succData := user

	c.IndentedJSON(http.StatusAccepted, responses.Success(message, succData))
}

func GenerateAccessTokenByUserID(c *gin.Context) {

	var req requests.TokenRequestByUserID
	if err := c.ShouldBindJSON(&req); err != nil {
		message := "payload error"
		errData := map[string]interface{}{
			"cause": err.Error(),
		}

		c.IndentedJSON(http.StatusBadRequest, responses.Error(message, errData))
		return
	}

	user, err := services.GetUser(req.UserID)
	if err != nil {
		message := "user not found"
		errData := map[string]interface{}{
			"cause": err.Error(),
		}

		c.IndentedJSON(http.StatusNotFound, responses.Error(message, errData))
		return
	}

	token, err := services.GetToken(user.UID)
	if err != nil {
		message := "error generating token"
		errData := map[string]interface{}{
			"cause": err.Error(),
		}

		c.IndentedJSON(http.StatusUnauthorized, responses.Error(message, errData))
		return
	}

	message := "access token was obtained successfully"
	succData := map[string]interface{}{
		"accessToken": token,
		"expiresIn":   time.Hour.Milliseconds(),
	}

	c.IndentedJSON(http.StatusAccepted, responses.Success(message, succData))
}
