package logger

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Logger represents a logger instance.
type Logger struct {
	logger *log.Logger
}

// NewLogger creates a new logger instance.
func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Println logs a message with a UUID prefix.
func (l *Logger) Println(c *gin.Context, v ...interface{}) {
	uuid, exists := c.Get("requestID")
	if !exists {
		return
	}
	uuidStr, ok := uuid.(string)
	if !ok {
		l.logger.Println("UUID is not a string")
		return
	}
	l.logger.Println(append([]interface{}{uuidStr}, v...)...)
}

// LoggerInstance is the global logger instance.
var LoggerInstance = NewLogger() // Initialize LoggerInstance

// Init sets the prefix for the global logger instance.
func Init() {
	LoggerInstance.logger.SetPrefix("APP: ") // Example of setting prefix
}

// To use this logger in other packages, import the logger package and use the LoggerInstance.
// Example: logger.LoggerInstance.Println(c, "uuid123", "Log message")
