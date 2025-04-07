package handlers

import (
	"backend/db"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const accessAlive = 100
const refreshAlive = 1000

var hmacSampleSecret = []byte(os.Getenv("JWTSECRET"))

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c echo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	username := user.Username
	password := user.Password
	database := db.Database{}.GetConnection()
	var isAuthenticated bool
	err = database.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2", username, password).Scan(&isAuthenticated)
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
			"exp":  time.Now().Unix() + refreshAlive,
			"name": username,
		})
		accessTokenString, err := accessToken.SignedString(hmacSampleSecret)
		if err != nil {
			panic(err)
		}
		refreshTokenString, _ := refreshToken.SignedString(hmacSampleSecret)
		fmt.Println(accessTokenString)
		return c.JSON(http.StatusOK, map[string]string{
			"accessToken":  accessTokenString,
			"refreshToken": refreshTokenString,
		})

	} else {
		return c.String(http.StatusUnauthorized, "Wrong creds")
	}
}

func refreshHandler(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "error while parsing token (maybe expired)",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["type"] != "refresh" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "error while parsing token (maybe expired)",
			})
		}
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"type": "access",
			"exp":  time.Now().Unix() + accessAlive,
			"name": claims["name"],
		})
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"type": "refresh",
			"exp":  time.Now().Unix() + refreshAlive,
			"name": claims["name"],
		})
		accessTokenString, err := accessToken.SignedString(hmacSampleSecret)
		if err != nil {
			panic(err)
		}
		refreshTokenString, _ := refreshToken.SignedString(hmacSampleSecret)
		fmt.Println(accessTokenString)
		return c.JSON(http.StatusOK, map[string]string{
			"accessToken":  accessTokenString,
			"refreshToken": refreshTokenString,
		})

	} else {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "error while parsing token (maybe expired)",
		})
	}
}
