package core

import (
	"blog_gin/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

func InitGorm() *os.File {
	level := strings.ToLower(global.Config.Mysql.LogLevel)
	LogLevel := logger.Info
	switch level {
	case "warn":
		LogLevel = logger.Warn
	case "error":
		LogLevel = logger.Error
	}

	// 打开一个文件，用于记录慢查询日志
	gormLogFile, err := os.OpenFile("log/gorm_slow_queries.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		global.Logrus.Fatal(err)
	}

	var out *log.Logger
	var conf logger.Config
	if level == "info" {
		out = log.New(os.Stdout, "\r\n", log.LstdFlags) // 使用标准输出
		conf = logger.Config{
			//SlowThreshold: time.Second, // 查询超过1秒钟将被视为慢查询
			SlowThreshold: 1 * time.Millisecond, // 查询超过1秒钟将被视为慢查询
			LogLevel:      LogLevel,             // 日志级别：Silent、Error、Warn、Info
			//IgnoreRecordNotFoundError: true,        // 忽略记录未找到的错误
			Colorful: true, // 开启彩色打印
		}
	} else {
		out = log.New(gormLogFile, "\r\n", log.LstdFlags) // 输出到日志文件
		conf = logger.Config{
			SlowThreshold: 1 * time.Millisecond, // 查询超过1秒钟将被视为慢查询
			LogLevel:      LogLevel,             // 日志级别：Silent、Error、Warn、Info
		}
	}

	//创建一个自定义的日志记录器
	newLogger := logger.New(
		out,
		conf,
	)

	// 连接数据库并配置GORM
	dsn := global.Config.Mysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 设置自定义日志记录器
	})

	if err != nil {
		global.Logrus.Fatal("无法连接数据库: " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		global.Logrus.Info("未知错误: " + err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	global.DB = db
	return gormLogFile
}
