package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotFound(c echo.Context) error {
	data := map[string]interface{}{
		"message": "Route not found",
		"error":   true,
		"status":  http.StatusNotFound,
		"home":    "/api/v1/",
	}
	return c.JSONPretty(http.StatusNotFound, data, "  ")
}
