package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/tsaridas/gofun/internal/routes"
	"github.com/tsaridas/gofun/pkg/models"
)

func main() {
	models.InitDB("gofun.db")
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	routes.SetupRoutes(r)

	r.Run(":3000")
}
