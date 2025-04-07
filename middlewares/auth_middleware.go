package middlewares

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var hmacSampleSecret = []byte(os.Getenv("JWTSECRET"))

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
			c.Set("username", claims["name"])
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "error while parsing token (maybe expired)",
			})
		}
		return next(c)
	}
}
