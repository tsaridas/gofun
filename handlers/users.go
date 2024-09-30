package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tsaridas/gofun/logger"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make([]User, 50) // Increased the number of users

func init() {
	for i := 0; i < len(users); i++ { // Dynamically adjust the loop based on the length of users
		users[i] = User{Name: "User" + strconv.Itoa(i), Age: rand.Intn(100)}
	}
}

func GetUsers(c *gin.Context) {

	pageStr := c.Query("page")
	searchQuery := c.Query("search")
	logger.LoggerInstance.Println(c, "PageStr is:", string(pageStr))         // Use the logger instance
	logger.LoggerInstance.Println(c, "searchQuery is:", string(searchQuery)) // Use the logger instance

	var page int
	if pageStr == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr)
	}
	start := (page - 1) * 10
	end := start + 10

	if end > len(users) { // Dynamically adjust the condition to match the current number of users
		end = len(users)
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

	// Filter users based on search query
	var filteredUsers []User
	for _, user := range users[start:end] {
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(searchQuery)) {
			filteredUsers = append(filteredUsers, user)
		}
	}

	// Calculate total pages
	totalPages := len(users) / 10
	if len(users)%10 > 0 {
		totalPages++
	}

	if page > totalPages {
		c.JSON(http.StatusNotFound, gin.H{"error": "No more pages available"})
		return
	}
	logger.LoggerInstance.Println(c, "Sending results to client is:", filteredUsers) // Use the logger instance

	c.JSON(http.StatusOK, gin.H{"users": filteredUsers, "page": page, "total_pages": totalPages})
}
