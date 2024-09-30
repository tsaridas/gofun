package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsaridas/gofun/internal/handlers"
	"github.com/tsaridas/gofun/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middleware.UUIDMiddleware())
	api := r.Group("/api")
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
