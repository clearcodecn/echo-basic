package middleware

import (
	"echo-basic/app/model"
	repo2 "echo-basic/app/repo"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type MyCtx struct {
	echo.Context
	user *model.User
}

func WrapJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		myctx := &MyCtx{
			Context: ctx,
			user:    nil,
		}

		user, ok := ctx.Get("user").(*jwt.Token)
		if !ok {
			return echo.ErrUnauthorized
		}

		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)

		u, err := repo2.UserRepo.FindByName(name)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return echo.ErrUnauthorized
			}
			return err
		}
		myctx.user = u

		return next(myctx)
	}
}

func (ctx *MyCtx) Uid() int {
	return ctx.user.ID
}

func (ctx *MyCtx) UserInfo() *model.User {
	return ctx.user
}
