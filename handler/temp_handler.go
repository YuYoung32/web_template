package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"web_template/database/mysql"
	project_redis "web_template/database/redis"
	"web_template/log"
)

func init() {
	// TODO
}

func PingHandler(ctx *gin.Context) {
	logger := log.GetLogger() // 惯用法，获取具有当前函数调用栈的日志实例
	mysqlDB := mysql.GetDBInstance()
	redisDB := project_redis.GetRedisInstance()

	var m mysql.TempModel
	err := mysqlDB.Table("temp").Where("id = ?", 1).First(&m)
	if err != nil {
		if err.Error == gorm.ErrRecordNotFound {
			// 无事发生
		} else {
			logger.Error("mysql error: ", err)
			ctx.Abort()
			return
		}
	}
	redisDB.Set(context.Background(), "test", m.Name, 0)

	logger.Debug("ping")
	ctx.String(200, "pong")
}

func AuthPingHandler(ctx *gin.Context) {
	ctx.String(200, "auth ping")

}
