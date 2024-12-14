package handlers

import (
	"github.com/labstack/echo/v4"
)

func getFile(c echo.Context) error {
	return c.File("resources/fells_loop.gpx")
}
