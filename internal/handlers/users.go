package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tsaridas/gofun/internal/logger"
)

var userLog = logger.NewLogger()

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make([]User, 50)

func init() {
	for i := 0; i < len(users); i++ {
		users[i] = User{Name: "User" + strconv.Itoa(i), Age: rand.Intn(100)}
	}
}

func GetUsers(c *gin.Context) {

	pageStr := c.Query("page")
	searchQuery := c.Query("search")
	userLog.LogRequest(c, "PageStr is: %s", string(pageStr))
	userLog.LogRequest(c, "SearchQuery is: %s", string(searchQuery))

	var page int
	if pageStr == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr)
	}
	start := (page - 1) * 10
	end := start + 10

	if end > len(users) {
		end = len(users)
	}

	if start < 0 {
		start = 0
	}
	if end > len(users) {
		end = len(users)
	}
	if start > end {
		start = end
	}

	var filteredUsers []User
	for _, user := range users[start:end] {
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(searchQuery)) {
			filteredUsers = append(filteredUsers, user)
		}
	}

	totalPages := len(users) / 10
	if len(users)%10 > 0 {
		totalPages++
	}

	if page > totalPages {
		c.JSON(http.StatusNotFound, gin.H{"error": "No more pages available"})
		return
	}
	userLog.LogRequest(c, "Sending results to client is: %v", filteredUsers)

	c.JSON(http.StatusOK, gin.H{"users": filteredUsers, "page": page, "total_pages": totalPages})
}
