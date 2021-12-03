package main

import (
	"encoding/json"
	"io/ioutil"

	"git.nextchat.org/nextchat/nextchat-go/controllers"
	"git.nextchat.org/nextchat/nextchat-go/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Start the new go echo app
	echo.NotFoundHandler = controllers.NotFound
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
	g.GET("/", controllers.HomeHandler)
	// Get the routes
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	utils.CheckError(err)
	ioutil.WriteFile("routes.json", data, 0o644)
	// Start app
	e.Logger.Fatal(e.Start(":8080"))
}
