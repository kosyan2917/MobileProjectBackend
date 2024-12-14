package main

import (
	"backend/handlers"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	handlers.SetupRoutes(e)
	fmt.Println(*e.Routes()[0])
	e.Logger.Fatal(e.Start(":8080"))

}
