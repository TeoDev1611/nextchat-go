package controllers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"git.nextchat.org/nextchat/nextchat-go/database"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"

	"git.nextchat.org/nextchat/nextchat-go/models"
	"git.nextchat.org/nextchat/nextchat-go/security"
	"github.com/google/uuid"
)

var (
	usersCollection *mongo.Collection = database.GetCollection("users")
	newUserModel    models.NewUser
	reqBody         models.CreateUserData
	ctx, cancel     = context.WithTimeout(context.Background(), 100*time.Second)
)

func CreateAccount(c *gin.Context) {
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "The request was invalid you need pass the correct model",
			"status":  http.StatusUnprocessableEntity,
		})
		return
	}

	// Check the user data is empty
	if reqBody.Username == " " && len(reqBody.Username) < 4 && len(reqBody.Username) > 15 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "The username is not correct, minimum 4 chars and max 15 chars.",
			"status":  http.StatusBadRequest,
		})
		return
	}

	u, errUrl := url.Parse(reqBody.ProfileImage)
	if errUrl != nil {
		fmt.Print("------------------ ERROR -------------------")
		log.Error(errUrl.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Internal error in the url parsing",
			"status":  http.StatusInternalServerError,
		})
		return
	} else if u.Host == " " || u.Scheme == " " {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      true,
			"message":    "Pass the correct url format for the profile image.",
			"status":     http.StatusBadRequest,
			"suggestion": "Check the url format",
		})
		return
	}

	/**  Generate the Data for the requests **/

	// Generate a ID
	if newUserModel.ID.String() == " " {
		newUserModel.ID = uuid.New()
	}

	// Add the user and the password profile image
	newUserModel.Username = reqBody.Username
	newUserModel.Password = reqBody.Password
	newUserModel.ProfileImage = reqBody.ProfileImage

	// Generate the Time
	if newUserModel.JoinedAt.String() == " " {
		newUserModel.JoinedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	}

	// Generate the Recovers code
	if len(newUserModel.RecoverCodes) == 0 {
		newUserModel.RecoverCodes = security.GenerateCodes()
	}

	// TODO: Check if is correct encripted
	// Generate the password encryption
	if newUserModel.Password != " " && len(newUserModel.Password) < 8 {
		encryption, ok := security.EncryptArgon(newUserModel.Password)
		if ok {
			newUserModel.Password = encryption
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":      true,
				"message":    "Error encrypting the password :{",
				"status":     http.StatusInternalServerError,
				"suggestion": "Check the encryption algorithm",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      true,
			"message":    "The password is too short, must be 8 characters.",
			"status":     http.StatusBadRequest,
			"suggestion": "Check the length of the password must be 8 chars :)",
		})
		return
	}

	result, insertErr := usersCollection.InsertOne(ctx, newUserModel)
	if insertErr != nil {
		fmt.Print("------------------ ERROR -------------------")
		log.Error(insertErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Route not found",
			"error":      true,
			"status":     http.StatusInternalServerError,
			"suggestion": "Check if the database is working correctly",
		})
		return
	}

	defer cancel()

	// Return the data
	c.JSON(http.StatusCreated, gin.H{
		"message":      "User Created",
		"error":        false,
		"status":       http.StatusCreated,
		"data":         newUserModel,
		"databaseInfo": result,
	})
}
