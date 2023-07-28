package main

import (
	"go-games-api/controllers"
	"go-games-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironment()
	initializers.DatabaseConnect()
}

func main() {
	router := gin.Default()
	api := router.Group("/api")

	// word controller
	api.GET("/word/:id", controllers.WordById)
	api.POST("/word/random", controllers.WordRandom)

	router.Run()
}
