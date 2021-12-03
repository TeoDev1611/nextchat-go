package main

import (
	"git.nextchat.org/nextchat/nextchat-go/routes"
	"git.nextchat.org/nextchat/nextchat-go/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := routes.Router()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "POST", "GET", "DELETE"},
	}))
	if utils.GetEnv("DEPLOY") == "on" {
		gin.SetMode(gin.ReleaseMode)
		r.Run()
	} else {
		r.Run(":3000")
	}
}
