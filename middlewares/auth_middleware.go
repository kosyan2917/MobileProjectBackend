package middlewares

import (
	"backend/db"
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Получаем токен из заголовка Authorization
		token := c.Request().Header.Get("Authorization")
		database := db.Database{}.GetConnection()
		var username string
		err := database.QueryRow("SELECT username from users where token = $1", token).Scan(&username)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.Set("username", "")
				return next(c)
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "error while extracting token",
				})
			}
		}
		c.Set("username", username)
		return next(c)
	}
}
