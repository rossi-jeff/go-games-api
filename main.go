package main

import (
	"go-games-api/controllers"
	"go-games-api/initializers"

	_ "go-games-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvironment()
	initializers.DatabaseConnect()
}

// @Title Games API in Go
// @Description backend implementation to match typescript version
// @version         0.1.0
// @contact.name   Jeff Rossi
// @contact.url    https://resume-svelte.jeff-rossi.com/
// @contact.email  inquiries@jeff-rossi.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	router := gin.Default()

	// docs route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	// auth controller
	api.POST("/auth/register", controllers.Register)
	api.POST("/auth/login", controllers.Login)
	// word controller
	api.GET("/word/:id", controllers.WordById)
	api.POST("/word/random", controllers.WordRandom)
	// code breaker controller
	api.GET("/code_breaker", controllers.CodeBreakerIndex)
	api.GET("/code_breaker/:id", controllers.CodeBreakerById)
	// concentration controller
	api.GET("/concentration", controllers.ConcentrationIndex)
	api.GET("/concentration/:id", controllers.ConcentrationById)
	// free cell controller
	api.GET("/free_cell", controllers.FreeCellIndex)
	api.GET("/free_cell/:id", controllers.FreeCellById)
	// guess word controller
	api.GET("/guess_word", controllers.GuessWordIndex)
	api.GET("/guess_word/:id", controllers.GuessWordById)
	// hang man controller
	api.GET("/hang_man", controllers.HangManIndex)
	api.GET("/hang_man/:id", controllers.HangManById)
	// klondike controller
	api.GET("/klondike", controllers.KlondikeIndex)
	api.GET("/klondike/:id", controllers.KlondikeById)
	// poker square controller
	api.GET("/poker_square", controllers.PokerSquareIndex)
	api.GET("/poker_square/:id", controllers.PokerSquareById)
	// sea battle controller
	api.GET("/sea_battle", controllers.SeaBattleIndex)
	api.GET("/sea_battle/:id", controllers.SeaBattleById)
	// spider controller
	api.GET("/spider", controllers.SpiderIndex)
	api.GET("/spider/:id", controllers.SpiderById)
	// ten grand controller
	api.GET("/ten_grand", controllers.TenGrandIndex)
	api.GET("/ten_grand/:id", controllers.TenGrandById)
	// yacht controller
	api.GET("/yacht", controllers.YachtIndex)
	api.GET("/yacht/:id", controllers.YachtById)

	router.Run()
}
