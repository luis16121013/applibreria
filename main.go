package main

import (
	"github.com/labstack/echo/v4"
	"github.com/luis16121013/libreria/handlers"
	midd "github.com/luis16121013/libreria/middlewares"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.HelloHandler, midd.Test)
	e.GET("/users", handlers.GetUsers, midd.Authentication)
	e.Start(":3000")
}


