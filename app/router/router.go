package router

import (
	"echo-basic/app/controller"
	mymiddleware "echo-basic/app/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	idx = new(controller.Index)
)

func RouterRegister(e *echo.Echo) {
	e.GET("/", idx.Index)
	e.POST("/register", idx.ValidateDemo)
	e.POST("/login", idx.Login)

	e.Use(middleware.JWT([]byte("secret")))
	e.POST("/testjwt", idx.UserInfo, mymiddleware.WrapJWT)
}
