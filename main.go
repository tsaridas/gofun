package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/tsaridas/gofun/logger"
	"github.com/tsaridas/gofun/middleware"
	"github.com/tsaridas/gofun/routes"
)

func main() {
	logger.Init()
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	r.Use(middleware.UUIDMiddleware())

	routes.SetupRoutes(r)

	r.Run(":3000")
}
