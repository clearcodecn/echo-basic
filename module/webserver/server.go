package webserver

import (
	"context"
	mymiddleware "echo-basic/app/middleware"
	"echo-basic/app/router"
	"echo-basic/pkg/modules"
	"echo-basic/pkg/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"time"
)

var (
	webServer  *modules.Module
	echoGlobal *echo.Echo
)

func Start(addr string, autoTls bool, cert string, key string) {
	webServer = modules.Register("webServer", 1)

	echoGlobal = echo.New()
	echoGlobal.HideBanner = true
	echoGlobal.Debug = true
	echoGlobal.Validator = validator.Instance()

	// 错误处理.
	echoGlobal.HTTPErrorHandler = mymiddleware.ErrHandler
	// 全局中间件
	echoGlobal.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           nil,
		StackSize:         1024,
		DisableStackAll:   false,
		DisablePrintStack: false,
	}))
	// gzip 压缩.
	echoGlobal.Use(middleware.Gzip())
	echoGlobal.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: nil,
		Format: `{"time":"${time_rfc3339}","id":"${id}","method":"${method}","uri":"${uri}",` +
			`"status":${status},"bytes_in":${bytes_in},"bytes_out":${bytes_out},"remote_ip":"${remote_ip}"}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
		Output:           nil,
	}))

	// 静态文件
	if staticdir, prefix := viper.GetString("server.staticDir"), viper.GetString("server.staticPrefix"); staticdir != "" && prefix != "" {
		echoGlobal.Static(prefix, staticdir)
	}

	router.RouterRegister(echoGlobal)

	var erch = make(chan error, 1)
	if autoTls {
		go func() {
			erch <- echoGlobal.StartAutoTLS(addr)
		}()
	} else if cert != "" && key != "" {
		go func() {
			erch <- echoGlobal.StartTLS(addr, cert, key)
		}()
	} else {
		go func() {
			erch <- echoGlobal.Start(addr)
		}()
	}

	for {
		select {
		case err := <-erch:
			log.Fatal(err)
		case <-webServer.Stop:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			echoGlobal.Shutdown(ctx)
			webServer.StopComplete()
		}
	}
}
