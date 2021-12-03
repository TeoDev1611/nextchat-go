package controllers

import (
	"net/http"
	"time"

	"git.nextchat.org/nextchat/nextchat-go/models"
	"git.nextchat.org/nextchat/nextchat-go/security"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateAccount(c echo.Context) error {
	var reqBody models.UserInfo
	if err := c.Bind(reqBody); err != nil {
		c.JSONPretty(http.StatusUnprocessableEntity, map[string]interface{}{
			"error":   true,
			"message": "The request was invalid you need pass the correct model",
			"status":  http.StatusUnprocessableEntity,
		}, "  ")
	}

	// Check the user data is empty
	if reqBody.Username == " " || reqBody.Password == " " {
		c.JSONPretty(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": "Fill all blanks!",
			"status":  http.StatusBadRequest,
		}, "  ")
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

	// TODO: Make better ok message and create the response correct
	return c.JSONPretty(http.StatusCreated, map[string]interface{}{
		"ok": true,
	}, "  ")
}
