package handlers

import (
	"backend/db"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const accessAlive = 100
const refreshAlive = 1000

var hmacSampleSecret = os.Getenv("JWTSECRET")

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
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"type": "access",
			"exp":  time.Now().Unix() + accessAlive,
			"name": username,
		})
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"type": "refresh",
			"exp":  time.Now().Unix() + accessAlive,
			"name": username,
		})
		accessTokenString, _ := accessToken.SignedString(hmacSampleSecret)
		refreshTokenString, _ := refreshToken.SignedString(hmacSampleSecret)

		return c.JSON(http.StatusOK, map[string]string{
			"accessToken":  accessTokenString,
			"refreshToken": refreshTokenString,
		})

	} else {
		return c.String(http.StatusUnauthorized, "Wrong creds")
	}
}
