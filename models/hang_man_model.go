package models

import "go-games-api/enum"

type HangMan struct {
	BaseModel
	Correct, Wrong string
	Status         enum.GameStatus
	Score          int
	UserId         int64 `json:"user_id"`
	User           User  `json:"user,omitempty"`
	WordId         int64 `json:"word_id"`
	Word           Word  `json:"word,omitempty"`
}

type HangManPaginated struct {
	Items []HangMan
	Paginated
}
