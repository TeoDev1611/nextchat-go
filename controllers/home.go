package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	data := map[string]string{
		"message": "Welcome to NextChat!",
		"status":  "Ok!",
	}
	return c.JSONPretty(http.StatusOK, data, "  ")
}
