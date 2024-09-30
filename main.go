package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tsaridas/gofun/routes"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	routes.SetupRoutes(r)

	r.Run(":3000")
}
