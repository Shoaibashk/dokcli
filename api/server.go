package api

import (
	"github.com/labstack/echo/v4"
)

func Server() {
	// Echo instance
	e := echo.New()
	e.HideBanner = true
	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Route => handler
	// e.GET("/", handler.Root)

	// e.GET("/users/:id", handler.GetUser)

	// e.POST("/users", handler.AddCat)
	// Start server
	e.Static("/", "/ui/dist")
	e.Logger.Fatal(e.Start(":1212"))
}
