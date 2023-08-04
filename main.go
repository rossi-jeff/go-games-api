package main

import (
	"go-games-api/controllers"
	"go-games-api/initializers"

	"github.com/gin-contrib/cors"

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

	// cors default
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

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
	api.GET("/code_breaker/progress", controllers.CodeBreakerInProgress)
	api.GET("/code_breaker/:id", controllers.CodeBreakerById)
	api.POST("/code_breaker", controllers.CodeBreakerCreate)
	api.POST("/code_breaker/:id/guess", controllers.CodeBreakerGuess)
	// concentration controller
	api.GET("/concentration", controllers.ConcentrationIndex)
	api.GET("/concentration/:id", controllers.ConcentrationById)
	api.POST("/concentration", controllers.ConcentrationCreate)
	api.PATCH("/concentration/:id", controllers.ConcentrationUpdate)
	// free cell controller
	api.GET("/free_cell", controllers.FreeCellIndex)
	api.GET("/free_cell/:id", controllers.FreeCellById)
	api.POST("/free_cell", controllers.FreeCellCreate)
	api.PATCH("/free_cell/:id", controllers.FreeCellUpdate)
	// guess word controller
	api.GET("/guess_word", controllers.GuessWordIndex)
	api.GET("/guess_word/progress", controllers.GuessWordInProgress)
	api.GET("/guess_word/:id", controllers.GuessWordById)
	api.POST("/guess_word", controllers.GuessWordCreate)
	api.POST("/guess_word/:id/guess", controllers.GuessWordGuess)
	api.POST("/guess_word/hint", controllers.GuessWordHints)
	// hang man controller
	api.GET("/hang_man", controllers.HangManIndex)
	api.GET("/hang_man/progress", controllers.HangManInProgress)
	api.GET("/hang_man/:id", controllers.HangManById)
	api.POST("/hang_man", controllers.HangManCreate)
	api.POST("/hang_man/:id/guess", controllers.HangManGuess)
	// klondike controller
	api.GET("/klondike", controllers.KlondikeIndex)
	api.GET("/klondike/:id", controllers.KlondikeById)
	api.POST("/klondike", controllers.KlondikeCreate)
	api.PATCH("/klondike/:id", controllers.KlondikeUpdate)
	// poker square controller
	api.GET("/poker_square", controllers.PokerSquareIndex)
	api.GET("/poker_square/:id", controllers.PokerSquareById)
	api.POST("/poker_square", controllers.PokerSquareCreate)
	api.PATCH("/poker_square/:id", controllers.PokerSquareUpdate)
	// sea battle controller
	api.GET("/sea_battle", controllers.SeaBattleIndex)
	api.GET("/sea_battle/progress", controllers.SeaBattleInProgress)
	api.GET("/sea_battle/:id", controllers.SeaBattleById)
	api.POST("/sea_battle", controllers.SeaBattleCreate)
	api.POST("/sea_battle/:id/ship", controllers.SeaBattleShip)
	api.POST("/sea_battle/:id/fire", controllers.SeaBattleFire)
	// spider controller
	api.GET("/spider", controllers.SpiderIndex)
	api.GET("/spider/:id", controllers.SpiderById)
	api.POST("/spider", controllers.SpiderCreate)
	api.PATCH("/spider/:id", controllers.SpiderUpdate)
	// ten grand controller
	api.GET("/ten_grand", controllers.TenGrandIndex)
	api.GET("/ten_grand/progress", controllers.TenGrandInProgress)
	api.GET("/ten_grand/:id", controllers.TenGrandById)
	api.POST("/ten_grand", controllers.TenGrandCreate)
	api.POST("/ten_grand/:id/roll", controllers.TenGrandRoll)
	api.POST("/ten_grand/options", controllers.TenGrandOptions)
	api.POST("/ten_grand/:id/score", controllers.TenGrandScore)
	// yacht controller
	api.GET("/yacht", controllers.YachtIndex)
	api.GET("/yacht/progress", controllers.YachtInProgress)
	api.GET("/yacht/:id", controllers.YachtById)
	api.POST("/yacht", controllers.YachtCreate)
	api.POST("/yacht/:id/roll", controllers.YachtRoll)
	api.POST("/yacht/:id/score", controllers.YachtScore)

	router.Run()
}
