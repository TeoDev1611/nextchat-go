package main

import (
	"git.nextchat.org/nextchat/nextchat-go/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Start the new go echo app
	e := echo.New()
	// Setup middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ ${method} ]: ${uri} -> ${status} || LATENCY: ${latency} |> HOST: ${host}\n",
	}))
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Setup routes
	g := e.Group("/api/v1")
	g.GET("/", routes.HomeHandler)
	// Start app
	e.Logger.Fatal(e.Start(":8080"))
}
