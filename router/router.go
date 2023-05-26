package router

import (
	"github.com/gin-gonic/gin"
	"web_template/handler"
	"web_template/middleware"
)

// PingRouter ping路由 路由分组更清晰
func PingRouter(engine *gin.Engine) {
	group := engine.Group("/ping")
	group.POST("/ping", handler.PingHandler)                               // 匹配 POST /ping/ping
	group.GET("/ping", middleware.AuthMiddleware, handler.AuthPingHandler) // 匹配 POST /ping/auth
}

func LoadAllRouter(engine *gin.Engine) {
	PingRouter(engine)
}
