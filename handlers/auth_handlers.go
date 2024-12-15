package handlers

import (
	"backend/db"
	"backend/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	database := db.Database{}.GetConnection()
	var isAuthenticated bool
	err := database.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2", username, password).Scan(&isAuthenticated)
	if err != nil {
		panic(err)
	}
	if isAuthenticated {
		token := utils.RandStringRunes(16)
		database.Exec("UPDATE users SET token = $1, started = $2 where username = $3", token, time.Now(), username)
		return c.JSONBlob(http.StatusOK, []byte(fmt.Sprintf(`{"message": "success", "token": "%s"}`, token)))
	} else {
		return c.String(http.StatusUnauthorized, "Wrong creds")
	}
}
