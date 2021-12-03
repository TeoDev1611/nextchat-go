package database

import (
	"context"
	"time"

	"git.nextchat.org/nextchat/nextchat-go/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(collection string) *mongo.Collection {
	URI := utils.GetEnv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))

	utils.CheckError(err)

	database := client.Database("nextchat-api")

	return database.Collection(collection)
}
