package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UUIDMiddleware creates a middleware that logs requests with a UUID.
func UUIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()
		c.Set("requestID", reqID) // Set the request ID in the context

		// Start time
		start := time.Now()

		// Process request
		c.Next()

		// Log the request with UUID
		logRequest(c, reqID, start)
	}
}

// logRequest logs the request details including the UUID.
func logRequest(c *gin.Context, reqID string, start time.Time) {
	duration := time.Since(start)
	status := c.Writer.Status()
	method := c.Request.Method
	path := c.Request.URL.Path
	clientIP := c.ClientIP()

	// Custom log format
	logMessage := " %s | %d | %s | %s | %s | %s | Request ID: %s"
	log.Printf(logMessage, time.Now().Format("2006/01/02 - 15:04:05"), status, duration, clientIP, method, path, reqID)
}
