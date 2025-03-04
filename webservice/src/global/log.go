package global

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type _JunuoLogFormatter struct {
}

func (f *_JunuoLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buf bytes.Buffer
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] %s in %s:%v - %s\n", timestamp, entry.Level.String(), entry.Caller.Function, entry.Caller.Line, entry.Message)
	buf.WriteString(logMsg)
	return buf.Bytes(), nil
}

func initLogSettings() {
	logrus.SetFormatter(&_JunuoLogFormatter{})
	logrus.SetReportCaller(true)
	path := "logs/webservice-"
	writer, err := rotatelogs.New(
		path+"%Y-%m-%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.Fatalf("Failed to initialize log file rotator: %v", err)
	}
	logrus.SetOutput(writer)
}

func JunuoWebLogMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		logrus.Infof("statusCode: %d, latency: %d, clientIP: %s, method: %s, path: %s", context.Writer.Status(), latency, context.ClientIP(), context.Request.Method, context.Request.URL.Path)
	}
}
