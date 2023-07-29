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
	// code breaker controller
	api.GET("/code_breaker", controllers.CodeBreakerIndex)
	api.GET("/code_breaker/:id", controllers.CodeBreakerById)
	// concentration controller
	api.GET("/concentration", controllers.ConcentrationIndex)
	api.GET("/concentration/:id", controllers.ConcentrationById)

	router.Run()
}
