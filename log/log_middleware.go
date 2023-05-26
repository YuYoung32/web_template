package log

import (
	"github.com/gin-gonic/gin"
	"time"
)

// Middleware 日志中间件
func Middleware(ctx *gin.Context) {
	// 开始时间
	startTime := time.Now()

	// 处理请求
	ctx.Next()

	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 请求方式
	reqMethod := ctx.Request.Method

	// 请求路由
	reqUri := ctx.Request.RequestURI

	// 状态码
	statusCode := ctx.Writer.Status()

	// 请求IP
	clientIP := ctx.ClientIP()

	// 日志格式
	MainLogger.Debugf("| %3d | %13v | %15s | %s | %s |",
		statusCode,
		latencyTime,
		clientIP,
		reqMethod,
		reqUri,
	)
}
