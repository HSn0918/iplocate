package utils

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	// Log 是全局日志实例
	Log *logrus.Logger
)

// 初始化日志
func init() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
}

// SetupLogger 配置日志
func SetupLogger(debug bool, logFile string) {
	if debug {
		Log.SetLevel(logrus.DebugLevel)
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
			ForceColors:     true,
		})
		Log.Debug("调试模式已启用")
	} else {
		Log.SetLevel(logrus.InfoLevel)
	}

	// 如果指定了日志文件，则同时输出到文件和标准输出
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			// 同时输出到文件和标准输出
			mw := io.MultiWriter(os.Stdout, file)
			Log.SetOutput(mw)
			Log.Infof("日志将同时写入文件: %s", logFile)
		} else {
			Log.Warnf("无法打开日志文件: %v", err)
		}
	}
}

// HTTPLogger 记录HTTP请求的日志
type HTTPLogger struct {
	Debug bool
}

// NewHTTPLogger 创建一个新的HTTP日志记录器
func NewHTTPLogger(debug bool) *HTTPLogger {
	return &HTTPLogger{
		Debug: debug,
	}
}

// LogRequest 记录HTTP请求信息
func (l *HTTPLogger) LogRequest(method, url string, headers map[string][]string) {
	if l.Debug {
		Log.WithFields(logrus.Fields{
			"method":  method,
			"url":     url,
			"headers": headers,
		}).Debug("发送HTTP请求")
	}
}

// LogResponse 记录HTTP响应信息
func (l *HTTPLogger) LogResponse(statusCode int, responseTime time.Duration, responseSize int) {
	if l.Debug {
		Log.WithFields(logrus.Fields{
			"status_code":   statusCode,
			"response_time": responseTime,
			"size":          responseSize,
		}).Debug("收到HTTP响应")
	}
}

// LogError 记录HTTP错误信息
func (l *HTTPLogger) LogError(err error, method, url string) {
	Log.WithFields(logrus.Fields{
		"error":  err,
		"method": method,
		"url":    url,
	}).Error("HTTP请求失败")
}
