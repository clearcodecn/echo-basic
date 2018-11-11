package config

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var (
	redisClient redis.Conn
)

func InitRedis() error {
	c, err := redis.Dial("tcp", viper.GetString("redis.addr"))
	if err != nil {
		return err
	}
	redisClient = c
	return nil
}

func CloseRedis() {
	if redisClient != nil {
		redisClient.Close()
	}
}
