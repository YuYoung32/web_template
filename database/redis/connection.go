package project_redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"web_template/conf"
	"web_template/log"
)

var rdb *redis.Client

func init() {
	var err error
	moduleLogger := log.GetLogger()

	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.GlobalConfig.GetString("redis.host") + ":" + conf.GlobalConfig.GetString("redis.port"),
		Password: conf.GlobalConfig.GetString("redis.password"),
		DB:       conf.GlobalConfig.GetInt("redis.db"),
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		moduleLogger.Error("failed to connect redis: " + err.Error())
		panic("failed to connect redis")
	}
	moduleLogger.Info("connect redis success")
}

// GetRedisInstance 获取redis实例, 其他包使用
func GetRedisInstance() *redis.Client {
	if rdb == nil {
		log.GetLogger().Error("rdb is nil")
	}
	return rdb
}
