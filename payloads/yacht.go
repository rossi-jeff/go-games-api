package payloads

import (
	"go-games-api/enum"
	"go-games-api/models"
)

type YachtRollPayload struct {
	Keep []int
}

type YachtScorePayload struct {
	TurnId   int64
	Category enum.YachtCategoryString
}

type YachtRollResponse struct {
	Turn    models.YachtTurnJson
	Options []YachtScoreOption
}

type YachtScoreOption struct {
	Category enum.YachtCategoryString
	Score    int
}
