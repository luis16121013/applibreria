package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/luis16121013/libreria/response"
	"github.com/luis16121013/libreria/users"
)

func HelloHandler(c echo.Context) error {
	msg := response.NewMessage()
	return msg.SendData(c)
}

func GetUsers(c echo.Context) error {
	msg := response.NewMessage()
	userModel := users.NewUser()

	data, err := userModel.GetUsers()
	if err != nil {
		return msg.UsersNotFound(c)
	}
	return msg.SendUsers(c, data)
}
