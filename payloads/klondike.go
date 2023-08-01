package payloads

import "go-games-api/enum"

type KlondikeUpdatePayload struct {
	Status  enum.GameStatusString
	Moves   int
	Elapsed int
}
