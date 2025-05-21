package handlers

import (
	"backend/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type profile struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func getProfile(c echo.Context) error {
	var name string
	var image string
	username := c.Param("username")
	fmt.Println(username)
	database := db.Database{}.GetConnection()
	err := database.QueryRow("SELECT name, avatar FROM users WHERE username = $1", username).Scan(&name, &image)
	if err != nil {
		panic(err)
		return c.JSON(http.StatusNotFound, "")
	}
	res := profile{Name: name, Image: image}
	return c.JSON(http.StatusOK, res)
}
