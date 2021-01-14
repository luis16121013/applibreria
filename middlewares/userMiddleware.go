package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Test(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("test accept")
		next(c)
		return nil
	}
}

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("user pass")
		next(c)
		return nil
	}
}
