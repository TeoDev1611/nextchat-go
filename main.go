package main

import (
	"git.nextchat.org/nextchat/nextchat-go/routes"
	"git.nextchat.org/nextchat/nextchat-go/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := routes.Router()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "POST", "GET", "DELETE"},
	}))

	port := utils.GetEnv("PORT")

	if utils.GetEnv("DEPLOY") == "on" {
		gin.SetMode(gin.ReleaseMode)
		r.Run()
	} else {
		r.Run(":" + port)
	}
}
