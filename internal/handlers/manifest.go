package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetManifest(c *gin.Context) {
	message := "access API at " + time.Now().Format(time.RFC3339Nano)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
