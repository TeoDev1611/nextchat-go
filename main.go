package main

import (
	"git.nextchat.org/nextchat/nextchat-go/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ ${method} ]: ${uri} -> ${status} || LATENCY: ${latency_human} |> HOST: ${host}\n",
	}))
	e.GET("/", routes.HomeHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
