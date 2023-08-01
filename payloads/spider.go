package payloads

import "go-games-api/enum"

type SpiderCreatePayload struct {
	Suits enum.Suit
}

type SpiderUpdatePayload struct {
	Status  enum.GameStatusString
	Moves   int
	Elapsed int
}
