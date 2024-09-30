package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/tsaridas/gofun/internal/logger"
)

var wsLog = logger.NewLogger()

func WebSocket(c *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set websocket upgrade"})
		return
	}
	defer conn.Close()

	// Show info when client connects
	clientIP := c.ClientIP()
	userAgent := c.Request.Header.Get("User-Agent")
	connectMessage := gin.H{"type": "connect", "ip": clientIP, "userAgent": userAgent}
	connectJSON, _ := json.Marshal(connectMessage)
	conn.WriteMessage(websocket.TextMessage, connectJSON)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Generate random data
		randData := strings.Repeat("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", rand.Intn(5))
		randDataJSON, _ := json.Marshal(gin.H{"type": "randomData", "data": randData})
		err = conn.WriteMessage(websocket.TextMessage, randDataJSON)
		if err != nil {
			wsLog.LogRequest(c, "Failed to write message: %v", err) // Log when failed to write message
			return
		}
		wsLog.LogRequest(c, "Data sent to client: %v", randDataJSON) // Use the logger instance
	}
}
