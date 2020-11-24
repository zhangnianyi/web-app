package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

// 初始化连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
		viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),
		PoolSize:      viper.GetInt("redis.poolsize"),              // use default DB
	})

	_, err = rdb.Ping().Result()
	return err
}

func Close(){
	_ =rdb.Close()
}