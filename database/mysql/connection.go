package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
	"web_template/conf"
	"web_template/log"
)

var db *gorm.DB

func init() {
	var err error
	moduleLogger := log.GetLogger()

	var dsn = conf.GlobalConfig.GetString("mysql.username") + ":" +
		conf.GlobalConfig.GetString("mysql.password") + "@tcp(" +
		conf.GlobalConfig.GetString("mysql.host") + ":" +
		conf.GlobalConfig.GetString("mysql.port") + ")/" +
		conf.GlobalConfig.GetString("mysql.database")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(log.MainLogger, logger.Config{
			SlowThreshold:             0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Warn,
		}),
		PrepareStmt: true,
	})
	if err != nil {
		moduleLogger.Error("failed to connect database: " + err.Error())
		panic(err)
	}
	moduleLogger.Info("connect database success")
	sqlDB, err := db.DB()
	if err != nil {
		moduleLogger.Error("failed to get sqlDB: " + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 自动迁移
	autoMigrate()
}

// GetDBInstance 获取数据库实例
func GetDBInstance() *gorm.DB {
	return db
}

func autoMigrate() {
	db := GetDBInstance()
	autoMigrateTempModel(db)
}
