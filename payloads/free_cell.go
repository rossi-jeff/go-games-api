package payloads

import "go-games-api/enum"

type FreeCellUpdatePayload struct {
	Status  enum.GameStatusString
	Moves   int
	Elapsed int
}
