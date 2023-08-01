package payloads

import "go-games-api/enum"

type PokerSquareUpdatePayload struct {
	Status enum.GameStatusString
	Score  int
}
