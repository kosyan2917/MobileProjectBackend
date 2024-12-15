package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"os"
)

func listFiles(c echo.Context) error {
	username := c.Get("username")
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "you are not authorized",
		})
	}
	entries, _ := os.ReadDir(fmt.Sprintf("resources/%s", username))
	result := []string{}
	for _, entry := range entries {
		result = append(result, entry.Name())
	}

	return c.JSON(http.StatusOK, map[string][]string{
		"files": result,
	})
}

func getFile(c echo.Context) error {
	username := c.Get("username")
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "you are not authorized",
		})
	}
	file := c.Param("file")
	fmt.Println(file)
	decodedFile, _ := url.QueryUnescape(file)
	return c.File(fmt.Sprintf("resources/%s/%s", username, decodedFile))
}
