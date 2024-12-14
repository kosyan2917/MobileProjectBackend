package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(app *echo.Echo) {
	group := app.Group("/api")
	group.GET("/file", getFile)
}
