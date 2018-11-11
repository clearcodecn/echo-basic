package middleware

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"net/http"
)

func ErrHandler(err error, ctx echo.Context) {

	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	switch err.(type) {
	case *echo.HTTPError:
		he := err.(*echo.HTTPError)
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			msg = fmt.Sprintf("%v, %v", err, he.Internal)
		}
	case govalidator.Errors:
		code = http.StatusUnprocessableEntity
		msg = err.Error()
	default:
		if ctx.Echo().Debug {
			msg = err.Error()
		} else {
			msg = http.StatusText(code)
		}
	}

	if _, ok := msg.(string); ok {
		msg = echo.Map{"message": msg}
	}

	ctx.Echo().Logger.Error(err)

	// Send response
	if !ctx.Response().Committed {
		if ctx.Request().Method == echo.HEAD { // Issue #608
			err = ctx.NoContent(code)
		} else {
			err = ctx.JSON(code, msg)
		}
		if err != nil {
			ctx.Echo().Logger.Error(err)
		}
	}
}
