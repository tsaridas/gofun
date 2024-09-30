package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetManifest(c *gin.Context) {
	currentTime := time.Now()
	message := "access API at " + currentTime.Format(time.RFC3339Nano)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
