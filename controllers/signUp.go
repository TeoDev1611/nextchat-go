package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"git.nextchat.org/nextchat/nextchat-go/models"
	"git.nextchat.org/nextchat/nextchat-go/security"
	"github.com/google/uuid"
)

func CreateAccount(c *gin.Context) {
	var reqBody models.UserInfo
	if err := c.ShouldBindJSON(reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "The request was invalid you need pass the correct model",
			"status":  http.StatusUnprocessableEntity,
		})
	}

	// Check the user data is empty
	if reqBody.Username == " " || reqBody.Password == " " {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Fill all blanks!",
			"status":  http.StatusBadRequest,
		})
	}

	/**  Generate the ids **/

	// Generate a ID
	if reqBody.Id.String() == " " {
		reqBody.Id = uuid.New()
	}

	// Generate the Time
	if reqBody.JoinedAt.String() == " " {
		reqBody.JoinedAt = time.Now().UTC()
	}

	// Generate the Recovers code
	if len(reqBody.RecoverCodes) == 0 {
		reqBody.RecoverCodes = security.GenerateCodes()
	}

	// TODO: Check if is correct encripted
	// Generate the password encription
	if reqBody.Password != " " {
		encription, ok := security.EncryptArgon(reqBody.Password)
		if ok {
			reqBody.Password = encription
		}
	}
	c.JSON(http.StatusCreated, gin.H{
		"ok": true,
	})
}
