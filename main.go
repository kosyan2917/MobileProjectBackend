package main

import (
	"backend/db"
	"backend/handlers"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database := db.Database{}.GetConnection()
	err := database.Ping()
	if err != nil {
		panic(err)
	}
	handlers.SetupRoutes(e)
	fmt.Println(*e.Routes()[0])
	e.Logger.Fatal(e.Start(":1337"))
}
