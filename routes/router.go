package routes

import (
	"git.nextchat.org/nextchat/nextchat-go/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	home := router.Group("/api/nextchat/v1")
	home.GET("/", controllers.HomeHandler)
	home.POST("/user/create", controllers.CreateAccount)
	router.NoRoute(controllers.NotFound)
	return router
}
