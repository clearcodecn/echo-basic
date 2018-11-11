package main

import (
	"echo-basic/config"
	_ "echo-basic/config"
	"echo-basic/module/webserver"
	"echo-basic/module/wsserver"
	"echo-basic/pkg/modules"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"
)

var (
	rootCmd = cli.NewApp()
)

func init() {
	rootCmd.Name = "app"
	rootCmd.UsageText = "app"
	rootCmd.Commands = []cli.Command{
		server,
	}
	rootCmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config,c",
			Usage:  "config file name",
			EnvVar: "CONFIG_FILE",
		},
	}

	rootCmd.Before = func(ctx *cli.Context) error {

		file := ctx.String("c")

		if file != "" {
			viper.SetConfigFile(file)
		} else {
			// set Config
			viper.SetConfigName("config")
			viper.AddConfigPath("/etc/echo-basic")
			viper.AddConfigPath("$HOME/.echo-basic")
			viper.AddConfigPath(".")
		}

		go vipWatch()

		err := viper.ReadInConfig()

		if err != nil {
			return err
		}

		// connect to mysql
		if err = config.InitMysql(); err != nil {
			log.Error(err)
		}

		// connect to redis
		if err = config.InitRedis(); err != nil {
			log.Error(err)
		}

		return nil
	}

}

func main() {
	if err := rootCmd.Run(os.Args); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

var server = cli.Command{
	Name:      "server",
	ShortName: "s",
	Aliases:   nil,
	Usage:     "server [option]",
	UsageText: "server manage",
	Before:    nil,
	Action:    startServer,
}

func startServer(ctx *cli.Context) error {

	go webserver.Start(
		viper.GetString("server.addr"),
		viper.GetBool("server.autoTls"),
		viper.GetString("server.cert"),
		viper.GetString("server.key"),
	)

	// optional
	go wsserver.Start()

	ch := make(chan os.Signal)

	signal.Notify(ch,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL)

	select {
	case <-ch:

		config.CloseRedis()
		config.CloseMysql()

		modules.InitiateFullShutdown()
	case <-modules.GlobalShutdown:
	}
	// wait for shutdown to complete, panic after timeout
	time.Sleep(5 * time.Second)
	fmt.Println("===== TAKING TOO LONG FOR SHUTDOWN - PRINTING STACK TRACES =====")
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

	return errors.New("server has been shutdown!")

	return nil
}

// 更新配置
func vipWatch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		config.SetDebug(viper.GetBool("mysql.debug"))
	})
}
