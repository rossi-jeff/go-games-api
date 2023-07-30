package payloads

import "go-games-api/enum"

type ConcentrationUpdatePayload struct {
	Status  enum.GameStatus
	Moves   int
	Matched int
	Elapsed int
}
