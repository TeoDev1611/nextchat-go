package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":    "Route not found",
		"error":      true,
		"status":     http.StatusNotFound,
		"suggestion": "/api/nextchat/v1/",
	})
}
