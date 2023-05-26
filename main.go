package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_template/conf"
	"web_template/log"
	"web_template/middleware"
	"web_template/router"
)

func init() {
	// TODO 一些业务的初始化可以需要手动触发
}

func main() {
	logger := log.GetLogger()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// 模式设置
	if conf.GlobalConfig.GetString("mode") == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if conf.GlobalConfig.GetString("mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	// 中间件加载
	engine.Use(log.Middleware)
	engine.Use(middleware.CrosMiddleware)

	// 路由加载
	router.LoadAllRouter(engine)

	runAddr := conf.GlobalConfig.GetString("server.host") + ":" + conf.GlobalConfig.GetString("server.port")
	srv := &http.Server{
		Addr:    runAddr,
		Handler: engine,
	}
	go func() {
		logger.Info("server is listening on " + runAddr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Error(err)
			panic(err)
		}
	}()

	// 阻塞, 等待结束
	sig := <-sigCh
	logger.Info("receive signal: ", sig, ", start to exit...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(err)
	}
}
