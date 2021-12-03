package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to NextChat!",
		"status":  http.StatusOK,
		"error":   false,
	})
}
