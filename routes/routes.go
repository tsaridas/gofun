package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsaridas/gofun/handlers"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api") // {{ edit_1 }}
	{
		api.GET("/users", handlers.GetUsers)

	}
	manifest := r.Group("/")
	{
		manifest.GET("/manifest.json", handlers.GetManifest)
	}

	ws := r.Group("/ws")
	{
		ws.GET("/random", handlers.WebSocket)
	}

	// Organize routes into groups
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "No route found"})
	})
}
