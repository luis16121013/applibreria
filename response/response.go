package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ServiceResponseMSG interface {
	SendData(c echo.Context) error
}

//struct Response Messages
type ResponseMSG struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type Message struct {
	msg *ResponseMSG
}

//Create new Response Messages
func NewMessage() *Message {
	m := &ResponseMSG{}
	return &Message{m}
}

func (rm *ResponseMSG) defaultResponse() {
	rm.Status = http.StatusOK
	rm.Message = "success"
}

func (r *Message) SendData(c echo.Context) error {
	r.msg.defaultResponse()
	return c.JSON(http.StatusOK, r.msg)
}

func (r *Message) UsersNotFound(c echo.Context) error {
	r.msg.defaultResponse()
	r.msg.Message = "Users Not Found!"
	return c.JSON(http.StatusOK, r.msg)
}

func (r *Message) SendUsers(c echo.Context, data interface{}) error {
	r.msg.defaultResponse()
	r.msg.Message = "success"
	r.msg.Data = data
	return c.JSON(http.StatusOK, r.msg)
}
