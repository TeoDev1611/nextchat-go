package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotFound(c echo.Context) error {
	data := map[string]string{
		"message": "Route not found",
		"status":  "404",
		"home":    "/api/v1/",
	}
	return c.JSONPretty(http.StatusNotFound, data, "  ")
}
