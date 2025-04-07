package handlers

import (
	"backend/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	group := app.Group("/api")
	authGroup := group.Group("/auth")
	authGroup.POST("/login", loginHandler)
	authGroup.GET("/refresh", refreshHandler)

	filesGroup := group.Group("/tracks")
	filesGroup.Use(middlewares.AuthMiddleware)
	filesGroup.POST("/diff", getDiff)
	filesGroup.GET("/:file", getFile)
}
