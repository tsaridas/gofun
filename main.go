package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // Allow all origins
		// You can customize other options here
	}))
	r.Any("/*path", func(c *gin.Context) {
		message := "access API at " + time.Now().Format(time.RFC3339)
		c.JSON(http.StatusOK, gin.H{"message": message})
	})
	r.Run(":3000")
}
