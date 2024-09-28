package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make([]User, 20)

func init() {
	for i := 0; i < 20; i++ {
		users[i] = User{Name: "User" + strconv.Itoa(i), Age: rand.Intn(100)}
	}
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // Allow all origins
		// You can customize other options here
	}))
	r.GET("/api/users", func(c *gin.Context) {
		pageStr := c.Query("page")
		var page int
		if pageStr == "" {
			page = 1
		} else {
			page, _ = strconv.Atoi(pageStr)
		}
		start := (page - 1) * 10
		end := start + 10
		if end > 20 {
			end = 20
		}

		// Ensure start and end are within bounds and start is less than or equal to end
		if start < 0 {
			start = 0
		}
		if end > len(users) {
			end = len(users)
		}
		if start > end {
			start = end // Adjust start to be equal to end if it exceeds
		}

		// Calculate total pages
		totalPages := len(users) / 10
		if len(users)%10 > 0 {
			totalPages++
		}

		c.JSON(http.StatusOK, gin.H{"users": users[start:end], "page": page, "total_pages": totalPages})
	})
	r.Any("/manifest.json", func(c *gin.Context) {
		message := "access API at " + time.Now().Format(time.RFC3339)
		c.JSON(http.StatusOK, gin.H{"message": message})
	})
	r.Run(":3000")
}
