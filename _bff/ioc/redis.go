package ioc

import (
	"github.com/spf13/viper"
	"github.com/redis/go-redis/v9"
)

func InitRedis() redis.Cmdable {

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: "", // no password set
	})
	return rdb
}
