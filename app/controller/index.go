package controller

import (
	"echo-basic/app/httpio/in"
	"echo-basic/app/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Index struct{}

func (Index) Index(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{
		"hello": "world",
	})
}

func (idx Index) ValidateDemo(ctx echo.Context) error {
	req := new(in.ExamPostRegister)
	ctx.Bind(req)
	err := ctx.Validate(req)

	if err != nil {
		return err
	}

	return ctx.JSON(200, map[string]string{
		"hello": "world",
	})
}

func (idx Index) Login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if username == "user" && password == "password" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "user"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func (idx Index) UserInfo(ctx echo.Context) error {
	myctx := ctx.(*middleware.MyCtx)
	return ctx.JSON(200, map[string]interface{}{
		"userInfo": myctx.UserInfo(),
	})
}
