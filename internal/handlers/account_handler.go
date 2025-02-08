package handlers

import (
	"net/http"

	"pulsy/internal/responses"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var accountRequest struct {
		AccessName string `json:"accessName"`
		Password   string `json:"password"`
	}

	if err := c.ShouldBindJSON(&accountRequest); err != nil {
		message := "error Json"
		errData := map[string]interface{}{
			"cause": err.Error(),
		}

		c.IndentedJSON(http.StatusConflict, responses.Error(message, errData))
		return
	}

	/* services.InsertData("accounts", c.GetString("userID"),) */
}
