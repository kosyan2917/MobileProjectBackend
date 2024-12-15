package handlers

import (
	"backend/middlewares"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	group := app.Group("/api")
	group.POST("/auth", loginHandler)

	filesGroup := group.Group("/files")
	filesGroup.Use(middlewares.AuthMiddleware)
	filesGroup.GET("/", listFiles)
	filesGroup.GET("/:file", getFile)
}
