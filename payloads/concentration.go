package payloads

import "go-games-api/enum"

type ConcentrationUpdatePayload struct {
	Status  enum.GameStatusString
	Moves   int
	Matched int
	Elapsed int
}
