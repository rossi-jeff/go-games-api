package controllers

import (
	"database/sql"
	"go-games-api/enum"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"go-games-api/utilities"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// @Summary      Sea Battle
// @Description  paginated list of sea battle
// @Tags         Sea Battle
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.SeaBattlePaginatedJson
// @Router       /api/sea_battle [get]
func SeaBattleIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.SeaBattlePaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.SeaBattle{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Score DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Sea Battle by Id
// @Description  get a sea battle
// @Tags         Sea Battle
// @Accept       json
// @Produce      json
// @Param Id path int true "Sea Battle ID"
// @Success      200  {object} models.SeaBattleJson
// @Router       /api/sea_battle/{Id} [get]
func SeaBattleById(c *gin.Context) {
	// get id
	id := c.Param("id")

	seaBattle := models.SeaBattle{}
	initializers.DB.Preload("Ships.Points").Preload("Ships.Hits").Preload(clause.Associations).First(&seaBattle, id)

	// response
	c.JSON(http.StatusOK, seaBattle.Json())
}

// @Summary      Create Sea Battle
// @Description  create a sea battle
// @Tags         Sea Battle
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.SeaBattleCreatePayload		true	"Sea Battle Options"
// @Success      201  {object} models.SeaBattleJson
// @Router       /api/sea_battle [post]
func SeaBattleCreate(c *gin.Context) {
	params := payloads.SeaBattleCreatePayload{}
	userId := utilities.UserIdFromAuthHeader(c)

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)
	seaBattle := models.SeaBattle{}
	if userId > 0 {
		seaBattle.UserId = sql.NullInt64{Int64: int64(userId), Valid: true}
	}
	seaBattle.Axis = params.Axis
	seaBattle.CreatedAt = now
	seaBattle.UpdatedAt = now
	seaBattle.Status = enum.Playing
	initializers.DB.Save(&seaBattle)

	c.JSON(http.StatusCreated, seaBattle.Json())
}

// @Summary      Create Sea Battle Ship
// @Description  create a sea battle
// @Tags         Sea Battle
// @Accept       json
// @Produce      json
// @Param Id path int true "Sea Battle ID"
// @Param	data	body	payloads.SeaBattleShipPayload		true	"Sea Battle Options"
// @Success      201  {object} models.SeaBattleShipJson
// @Router       /api/sea_battle/{Id}/ship [post]
func SeaBattleShip(c *gin.Context) {
	params := payloads.SeaBattleShipPayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seaBattle := models.SeaBattle{}
	initializers.DB.Preload("Ships.Points").Preload(clause.Associations).First(&seaBattle, id)

	if params.Navy == enum.N0 {
		createPlayerShip(c, seaBattle, params.ShipType, params.Size, params.Points)
	} else {
		createOpponentShip(c, seaBattle, params.ShipType, params.Size)
	}
}

func createOpponentShip(c *gin.Context, seaBattle models.SeaBattle, shipType enum.ShipTypeString, size int) {
	points := getOpponentShipPoints(seaBattle, size)

	now := time.Now().Format(time.RFC3339)
	ship := models.SeaBattleShip{}
	ship.SeaBattleId = int64(seaBattle.Id)
	ship.CreatedAt = now
	ship.UpdatedAt = now
	ship.Type = enum.ShipType(enum.ShipTypeArrayIndex(string(shipType)))
	ship.Size = size
	ship.Navy = enum.Opponent
	initializers.DB.Save(&ship)

	for i := 0; i < len(points); i++ {
		point := points[i]
		gridPoint := models.SeaBattleShipGridPoint{}
		gridPoint.SeaBattleShipId = int64(ship.Id)
		gridPoint.CreatedAt = now
		gridPoint.UpdatedAt = now
		gridPoint.Horizontal = point.Horizontal
		gridPoint.Vertical = point.Vertical
		initializers.DB.Save(&gridPoint)
	}

	found := models.SeaBattleShip{}
	initializers.DB.Preload("Points").First(&found, ship.Id)

	c.JSON(http.StatusCreated, found.Json())
}

func getOpponentShipPoints(seaBattle models.SeaBattle, size int) []payloads.SeaBattlePoint {
	var points []payloads.SeaBattlePoint
	grid, H, V := emptyGrid(seaBattle.Axis)
	for i := 0; i < len(seaBattle.Ships); i++ {
		ship := seaBattle.Ships[i]
		if ship.Navy == enum.Opponent {
			for p := 0; p < len(ship.Points); p++ {
				point := ship.Points[p]
				grid[point.Horizontal][point.Vertical] = true
			}
		}
	}
	for len(points) < size {
		points = nil
		direction := payloads.Directions[rand.Intn(len(payloads.Directions))]
		idxH := rand.Intn(seaBattle.Axis)
		idxV := rand.Intn(seaBattle.Axis)
		count := 0
		for count < size {
			if idxH < 0 || idxH >= seaBattle.Axis || idxV < 0 || idxV >= seaBattle.Axis {
				break
			}
			Horizontal := H[idxH]
			Vertical := V[idxV]
			if grid[Horizontal][Vertical] {
				break
			}
			point := payloads.SeaBattlePoint{
				Horizontal: Horizontal,
				Vertical:   Vertical,
			}
			points = append(points, point)
			count++
			switch direction {
			case "N":
				idxV--
			case "S":
				idxV++
			case "E":
				idxH++
			case "W":
				idxH--
			}
		}
	}
	return points
}

func emptyGrid(axis int) (map[string]map[int]bool, []string, []int) {
	grid := map[string]map[int]bool{}
	H := payloads.HorizontalAxisMax[0:axis]
	V := payloads.VerticalAxisMax[0:axis]
	for h := 0; h < axis; h++ {
		inner := map[int]bool{}
		for v := 0; v < axis; v++ {
			inner[V[v]] = false
		}
		grid[H[h]] = inner
	}
	return grid, H, V
}

func createPlayerShip(c *gin.Context, seaBattle models.SeaBattle, shipType enum.ShipTypeString, size int, points []payloads.SeaBattlePoint) {
	now := time.Now().Format(time.RFC3339)
	ship := models.SeaBattleShip{}
	ship.SeaBattleId = int64(seaBattle.Id)
	ship.CreatedAt = now
	ship.UpdatedAt = now
	ship.Type = enum.ShipType(enum.ShipTypeArrayIndex(string(shipType)))
	ship.Size = size
	ship.Navy = enum.Player
	initializers.DB.Save(&ship)

	for i := 0; i < len(points); i++ {
		point := points[i]
		gridPoint := models.SeaBattleShipGridPoint{}
		gridPoint.SeaBattleShipId = int64(ship.Id)
		gridPoint.CreatedAt = now
		gridPoint.UpdatedAt = now
		gridPoint.Horizontal = point.Horizontal
		gridPoint.Vertical = point.Vertical
		initializers.DB.Save(&gridPoint)
	}

	found := models.SeaBattleShip{}
	initializers.DB.Preload("Points").First(&found, ship.Id)

	c.JSON(http.StatusCreated, found.Json())
}

// @Summary      Create Sea Battle Turn
// @Description  create a sea battle turn
// @Tags         Sea Battle
// @Accept       json
// @Produce      json
// @Param Id path int true "Sea Battle ID"
// @Param	data	body	payloads.SeaBattleFirePayload		true	"Sea Battle Options"
// @Success      201  {object} models.SeaBattleTurnJson
// @Router       /api/sea_battle/{Id}/fire [post]
func SeaBattleFire(c *gin.Context) {
	params := payloads.SeaBattleFirePayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seaBattle := models.SeaBattle{}
	initializers.DB.Preload("Ships.Points").Preload("Ships.Hits").Preload(clause.Associations).First(&seaBattle, id)

	if params.Navy == enum.N0 {
		playerTurn(c, seaBattle, params.Horizontal, params.Vertical)
	} else {
		opponentTurn(c, seaBattle)
	}
}

func playerTurn(c *gin.Context, seaBattle models.SeaBattle, Horizontal string, Vertical int) {
	target := enum.Miss
	shipType := -1
	now := time.Now().Format(time.RFC3339)

	for s := 0; s < len(seaBattle.Ships); s++ {
		ship := seaBattle.Ships[s]
		if ship.Navy == enum.Opponent {
			hits := len(ship.Hits)
			for p := 0; p < len(ship.Points); p++ {
				point := ship.Points[p]
				if point.Horizontal == Horizontal && point.Vertical == Vertical {
					if hits+1 == ship.Size {
						target = enum.Sunk
					} else {
						target = enum.Hit
					}
					shipType = ship.Type.EnumIndex()

					shipHit := models.SeaBattleShipHit{}
					shipHit.SeaBattleShipId = int64(ship.Id)
					shipHit.Horizontal = Horizontal
					shipHit.Vertical = Vertical
					shipHit.CreatedAt = now
					shipHit.UpdatedAt = now
					initializers.DB.Save(&shipHit)
					break
				}
			}
		}
		if target != enum.Miss {
			break
		}
	}

	turn := models.SeaBattleTurn{}
	turn.SeaBattleId = int64(seaBattle.Id)
	turn.CreatedAt = now
	turn.UpdatedAt = now
	turn.Target = target
	turn.Navy = enum.Player
	if shipType != -1 {
		turn.ShipType = enum.ShipType(shipType)
	}
	initializers.DB.Save(&turn)

	c.JSON(http.StatusCreated, turn.Json())
}

func opponentTurn(c *gin.Context, seaBattle models.SeaBattle) {
	target := enum.Miss
	shipType := -1
	now := time.Now().Format(time.RFC3339)
	Horizontal := ""
	Vertical := -1
	grid, H, V := emptyGrid(seaBattle.Axis)

	for t := 0; t < len(seaBattle.Turns); t++ {
		T := seaBattle.Turns[t]
		if T.Navy == enum.Opponent {
			grid[T.Horizontal][T.Vertical] = true
		}
	}

	for Horizontal == "" && Vertical == -1 {
		idxH := rand.Intn(seaBattle.Axis)
		idxV := rand.Intn(seaBattle.Axis)
		if !grid[H[idxH]][V[idxV]] {
			Horizontal = H[idxH]
			Vertical = V[idxV]
		}
	}

	for s := 0; s < len(seaBattle.Ships); s++ {
		ship := seaBattle.Ships[s]
		if ship.Navy == enum.Player {
			hits := len(ship.Hits)
			for p := 0; p < len(ship.Points); p++ {
				point := ship.Points[p]
				if point.Horizontal == Horizontal && point.Vertical == Vertical {
					if hits+1 == ship.Size {
						target = enum.Sunk
					} else {
						target = enum.Hit
					}
					shipType = ship.Type.EnumIndex()

					shipHit := models.SeaBattleShipHit{}
					shipHit.SeaBattleShipId = int64(ship.Id)
					shipHit.Horizontal = Horizontal
					shipHit.Vertical = Vertical
					shipHit.CreatedAt = now
					shipHit.UpdatedAt = now
					initializers.DB.Save(&shipHit)
					break
				}
			}
		}
		if target != enum.Miss {
			break
		}
	}

	turn := models.SeaBattleTurn{}
	turn.SeaBattleId = int64(seaBattle.Id)
	turn.CreatedAt = now
	turn.UpdatedAt = now
	turn.Target = target
	turn.Navy = enum.Player
	if shipType != -1 {
		turn.ShipType = enum.ShipType(shipType)
	}
	initializers.DB.Save(&turn)

	c.JSON(http.StatusCreated, turn.Json())
}

func SeaBattleInProgress(c *gin.Context) {
	userId := utilities.UserIdFromAuthHeader(c)
	seaBattlesJson := []models.SeaBattleJson{}
	if userId > 0 {
		seabattles := []models.SeaBattle{}
		initializers.DB.Where("user_id = ? AND Status = 1", userId).Preload(clause.Associations).Find(&seabattles)
		for i := 0; i < len(seabattles); i++ {
			seaBattle := seabattles[i].Json()
			seaBattlesJson = append(seaBattlesJson, seaBattle)
		}
	}
	c.JSON(http.StatusOK, seaBattlesJson)
}
