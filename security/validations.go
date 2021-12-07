package security

import (
	"context"
	"strings"

	"git.nextchat.org/nextchat/nextchat-go/database"
	"git.nextchat.org/nextchat/nextchat-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection = database.GetCollection("users")

// Check if any user is repeated from the users collection :p
func CheckRepeatedUser(user string) bool {
	var structureUsers *models.CreateUserData
	err := usersCollection.FindOne(context.TODO(), primitive.D{{Key: "username", Value: user}}).Decode(&structureUsers)
	if err != nil {
		return false
	}
	return strings.EqualFold(structureUsers.Username, user)
}

// Validate the password checker
func CheckSamePassword(username, password string) (bool, *models.CreateUserData) {
	var structureUsers *models.CreateUserData
	err := usersCollection.FindOne(context.TODO(), primitive.D{{Key: "username", Value: username}}).Decode(&structureUsers)
	if err != nil {
		return false, &models.CreateUserData{}
	}

	match := CheckEncript(password, structureUsers.Password)

	return match, structureUsers
}
