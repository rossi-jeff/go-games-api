package payloads

import "go-games-api/enum"

type TenGrandRollPayload struct {
	Quantity int
}

type TenGrandOptionsPayload struct {
	Dice []int
}

type TenGrandScoreOptions struct {
	TurnId  int
	Dice    []int
	Options []TenGrandOption
}

type TenGrandOption struct {
	Score    int
	Category enum.TenGrandCategoryString
}
