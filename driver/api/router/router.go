package router

import (
	"github.com/grandcolline/clean-arch-demo/driver/api/handler"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, handler handler.AppHandler) {
	e.POST("/users", handler.CreateUser)
	e.GET("/users", handler.GetUsers)
}
