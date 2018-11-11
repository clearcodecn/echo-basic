package config

import (
	"echo-basic/pkg/atomicbool"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

var engine *gorm.DB
var debug *atomicbool.AtomicBool

func init() {
	debug = atomicbool.NewBool(false)
}

func InitMysql() error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}

	debug.SetTo(viper.GetBool("mysql.debug"))

	if debug.IsSet() {
		db.LogMode(true)
	}

	engine = db

	return nil
}

func CloseMysql() {
	if engine != nil {
		engine.Close()
	}
}

func NewDb() *gorm.DB {
	return engine.New()
}

func SetDebug(b bool) {
	debug.SetTo(b)
	log.Warn("set mysql debug mode: ", b)
	if debug.IsSet() {
		engine.LogMode(true)
	} else {
		engine.LogMode(false)
	}
}
