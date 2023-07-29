package core

import (
	"blog_gin/global"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logFile *os.File

type CustomHook struct {
	file *os.File
}

// Levels 错误日志等级为以下的, 触发Fire
func (hook *CustomHook) Levels() []logrus.Level {
	return []logrus.Level{
		//logrus.InfoLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (hook *CustomHook) Fire(entry *logrus.Entry) error {
	// 输出到文件
	out := io.MultiWriter(hook.file)
	formatter := &logrus.JSONFormatter{}
	entry.Logger.SetFormatter(formatter)
	entry.Logger.SetOutput(out)
	return nil
}

// InitLogrus 初始化日志
func InitLogrus() {
	// 创建新的Logger实例
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetReportCaller(global.Config.ShowPassLine)

	formatter := &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:03:04",
		FullTimestamp:   true,
	}
	log.SetFormatter(formatter)
	// 设置日志输出的最低级别为DebugLevel

	logFile, err := os.OpenFile("log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("文件打开失败, err: %v", err)
	}

	// 开发环境才在控制台显示日志
	if global.Config.Server.Env != "dev" {
		// 创建并添加自定义Hook将Error及以上级别的日志输出到文件
		log.AddHook(&CustomHook{file: logFile})
		// 取消控制台输出
		log.SetOutput(io.Discard)
	}
	global.Logrus = log
}

func CloseLogFile() {
	// 在适当的时候调用此函数，手动关闭日志文件
	if logFile != nil {
		logFile.Close()
	}
}
