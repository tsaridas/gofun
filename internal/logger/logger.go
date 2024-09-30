package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return &Logger{logger}
}

func (l *Logger) LogRequest(c *gin.Context, logMessage string, v ...interface{}) {
	reqID, exists := c.Get("requestID")
	if exists {
		l.WithField("requestID", reqID).Infof(logMessage, v...)
	} else {
		l.WithField("requestID", "Not Found").Infof(logMessage, v...)
	}
}
