package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"web_template/conf"
)

var (
	// MainLogger 全局Logrus实例
	MainLogger = logrus.New()
)

// Init 配置Logrus
func init() {
	level := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}
	if conf.GlobalConfig.GetString("mode") == "dev" {
		MainLogger.SetLevel(level["debug"])
	} else {
		MainLogger.SetLevel(level[conf.GlobalConfig.GetString("log.level")])
	}

	MainLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     false,
	})

	// 日志文件存储
	logFilePath := conf.GlobalConfig.GetString("log.path") + "/" + "log.log"
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	MainLogger.SetOutput(io.MultiWriter(file, os.Stdout))

	GetLogger().Info("Logrus init success")
}

// GetLogger 获取有调用栈的日志实例, Field为获得调用方的包名函数名
func GetLogger() *logrus.Entry {
	return GetLoggerWithSkip(2)
}

// GetLoggerWithSkip 获取日志实例
//
// skip=1 Field为调用GetLogger的函数
//
// skip=2 Field为调用GetLogger的函数的上一级函数
//
// 以此类推
func GetLoggerWithSkip(skip int) *logrus.Entry {
	// 获取调用栈信息
	pc, _, _, _ := runtime.Caller(skip)
	// 获取函数名
	funcName := runtime.FuncForPC(pc).Name()
	return MainLogger.WithField("func", funcName)
}
